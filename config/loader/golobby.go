// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package loader

import (
	"log"

	cfg "github.com/chinmobi/gin-mvc/config"

	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
)

const CONFIG_TAG_NAME = "json"

func Load(cfg *cfg.Config, configsPath string) error {
	config, err := config.New(config.Options{
		Feeder: feeder.JsonDirectory{Path: configsPath},
	})
	if err != nil {
		return err
	}

	configServer(cfg, config)
	configLogger(cfg, config)
	configGrpool(cfg, config)

	return nil
}

func configServer(cfg *cfg.Config, config *config.Config) {
	key := "server"
	if config.AssignStruct(&cfg.Server, key, CONFIG_TAG_NAME) <= 0 {
		warnConfig(key)
	}
}

func configLogger(cfg *cfg.Config, config *config.Config) {
	key := "logger.file"
	if config.AssignStruct(&cfg.Logger.File, key, CONFIG_TAG_NAME) <= 0 {
		warnConfig(key)
	}
	key = "logger.console"
	if config.AssignStruct(&cfg.Logger.Console, key, CONFIG_TAG_NAME) <= 0 {
		warnConfig(key)
	}
}

func configGrpool(cfg *cfg.Config, config *config.Config) {
	key := "grpool"
	if config.AssignStruct(&cfg.Grpool, key, CONFIG_TAG_NAME) <= 0 {
		warnConfig(key)
	}
}

func warnConfig(key string) {
	log.Printf("[WARNING] No configurable for: %s\n", key)
}
