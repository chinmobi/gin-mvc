// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"os"
	"strings"

	"github.com/chinmobi/gin-mvc/app/ctx"
)

const (
	APP_HOME = "APP_HOME"
)

type Config struct {
	Logger Logger
	Server Server
	Grpool Grpool
}

func Default() *Config {
	config := &Config{
		Logger: Logger{
			File: File{
				Enabled: false,
				Level: "INFO",
				Filename: "APP_HOME/var/logs/ginmvc.log", // The APP_HOME will be resolved as real home path at runtime.
				MaxSize: 500, // megabytes
				MaxBackups: 3,
				MaxAge: 28, // days
				Compress: false,
			},
			Console: Console{
				Enabled: true,
				Level: "DEBUG",
			},
			Level: "INFO",
		},
		Server: Server{
			Port: normalizePort(os.Getenv(SERVER_PORT_ENV_NAME)),
		},
		Grpool: Grpool{
			Size: 8,
		},
	}

	return config
}

func (c *Config) ResolveWith(ctx *ctx.AppContext) {
	filename := c.Logger.File.Filename
	if !strings.HasPrefix(filename, APP_HOME) {
		return
	}

	prefix := len(APP_HOME)
	filename = filename[prefix:]

	filename = ctx.GetRealPath(filename)
	c.Logger.File.Filename = filename
}
