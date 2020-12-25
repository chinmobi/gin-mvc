// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"log"
	"os"
	"strconv"
)

const (
	SERVER_PORT_ENV_NAME = "PORT"
	DEFAULT_SERVER_PORT  = "8080"

	INVALID_SERVER_PORT_LOG_FMT = "Invalid server port env: [%s], using the default value: [%s]\n"
	EMPTY_SERVER_PORT_LOG_FMT   = "Null server port env [%s], using the default value: [%s]\n"
)

type Server struct {
	Port string
}

type Config struct {
	Server Server
}

func Default() *Config {
	config := &Config{
	}

	config.Server.Port = normalizePort(os.Getenv(SERVER_PORT_ENV_NAME))

	return config
}

func normalizePort(port string) string {
	if len(port) == 0 {
		log.Printf(EMPTY_SERVER_PORT_LOG_FMT, SERVER_PORT_ENV_NAME, DEFAULT_SERVER_PORT)
		return DEFAULT_SERVER_PORT
	}

	if i, err := strconv.Atoi(port); err == nil {
		if i > 0 {
			return port
		}
	}

	log.Printf(INVALID_SERVER_PORT_LOG_FMT, port, DEFAULT_SERVER_PORT)

	return DEFAULT_SERVER_PORT
}
