// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/security/web/auth"
)

func setUp(authConfig *auth.ProcessorConfigurer, app *app.App) {
	authConfig.AddAuthFailureFunc(RespAuthenticationError)

	// Setup the UserDetails service auth provider, to authenticate the UesrnamePasswordAuthToken
	authProvider := app.ServiceSupplier().GetAuthService().CreateAuthProvider()
	authConfig.AddProvider(authProvider)
}
