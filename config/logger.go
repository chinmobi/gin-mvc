// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

type File struct {
	Enabled    bool
	Level      string
	Filename   string
	MaxSize    uint32
	MaxBackups uint32
	MaxAge     uint32
	Compress   bool
}

type Console struct {
	Enabled    bool
	Level      string
}

type Logger struct {
	File    File
	Console Console
	Level   string
}
