// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware/auth"
	"github.com/chinmobi/gin-mvc/middleware/common"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security"
)

func Configure(config *mw.Configurer, authHandler security.AuthHandlerSetter, app *app.App) error {
	builder := config.Build(mw.MW_COMMON)
	if err := common.Configure(builder, app); err != nil {
		return err
	}

	builder = config.Build(mw.MW_AUTH)
	if err := auth.Configure(builder, authHandler, app); err != nil {
		return err
	}

	return nil
}
