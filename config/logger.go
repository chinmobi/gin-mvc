// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

type File struct {
	Enabled    bool   `json:"enabled"`
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxBackups int    `json:"max_backups"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
}

type Console struct {
	Enabled    bool   `json:"enabled"`
	Level      string `json:"level"`
}

type Logger struct {
	File    File      `json:"file"`
	Console Console   `json:"console"`
	Level   string    `json:"level"`
}
