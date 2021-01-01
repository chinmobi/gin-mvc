// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware/auth/httpbasic"
	"github.com/chinmobi/gin-mvc/middleware/auth/jwt"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security/auth"
)

func Configure(builder *mw.Builder, authConfig *auth.ProcessorConfigurer, app *app.App) error {
	setUp(authConfig)

	authGroup := auth.NewProcessorGroup(authConfig)

	// Configure each of needed auth processors
	// NOTE: The order of the processors is IMPORTANT!

	if err := jwt.Configure(authGroup, app); err != nil {
		return err
	}

	if err := httpbasic.Configure(authGroup, app); err != nil {
		return err
	}

	builder.AddMwAdapter(authGroup)

	return nil
}
