// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/chinmobi/gin-mvc/app"
	mwauth "github.com/chinmobi/gin-mvc/middleware/auth"
	"github.com/chinmobi/gin-mvc/middleware/common"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security/auth"
)

func Configure(config *mw.Configurer, authConfig *auth.ProcessorConfigurer, app *app.App) error {
	builder := config.Build(mw.MW_COMMON)
	if err := common.Configure(builder, app); err != nil {
		return err
	}

	builder = config.Build(mw.MW_AUTH)
	if err := mwauth.Configure(builder, authConfig, app); err != nil {
		return err
	}

	return nil
}
