// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"os"
	"path/filepath"
	"strings"
)

type AppContext struct {
	home string
}

func NewAppContext() *AppContext {
	ctx := &AppContext{
	}
	return ctx
}

func (ctx *AppContext) Init() error {
	path, err := guessHomePath()
	if err != nil {
		return err
	}

	if path == "" {
		exe, err := getExecutable()
		if err != nil {
			return err
		}

		path = filepath.Dir(exe)
	}

	ctx.home = path

	return nil
}

func (ctx *AppContext) HomePath() string {
	return ctx.home
}

func (ctx *AppContext) GetRealPath(relativePath string) string {
	return filepath.Join(ctx.home, relativePath)
}

func guessHomePath() (string, error) {
	exe, err := getExecutable()
	if err != nil {
		return "", err
	}

	base := filepath.Base(exe)
	if strings.HasSuffix(base, ".exe") {
		return resolveExePath(exe), nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return resolvePwdPath(pwd), nil
}

func getExecutable() (string, error) {
	exe, err := os.Executable()
	if err == nil {
		return exe, nil
	}
	if exe == "" {
		return "", err
	}

	return filepath.EvalSymlinks(exe)
}

func resolveExePath(path string) string {
	dir := filepath.Dir(path)
	if filepath.Base(dir) == "bin" {
		dir = filepath.Dir(dir)
	}
	return dir
}

func resolvePwdPath(pwd string) string {
	for !isRoot(pwd) {
		path := filepath.Join(pwd, "_deploy")
		if pathExists(path) {
			return path
		}
		pwd = filepath.Dir(pwd)
	}
	return ""
}

func isRoot(path string) bool {
	if path != "" {
		c := path[len(path)-1]
		return c == filepath.Separator
	}
	return true
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
