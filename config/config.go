// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"os"
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
				Enabled: true,
				Level: "INFO",
				Filename: "/tmp/ginmvc.log", // "/var/log/ginmvc/ginmvc.log"
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
