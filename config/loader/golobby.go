// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package loader

import (
	cfg "github.com/chinmobi/gin-mvc/config"

	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
)

func Load(cfg *cfg.Config, configsPath string) error {
	_, err := config.New(config.Options{
		Feeder: feeder.JsonDirectory{Path: configsPath},
	})
	if err != nil {
		return err
	}

	return nil
}
