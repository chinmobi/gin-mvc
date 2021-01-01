// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security/auth"
)

func Configure(builder *mw.Builder, authConfig *auth.ProcessorConfigurer, app *app.App) error {
	setUp(authConfig)

	authGrop := auth.NewProcessorGroup(authConfig)

	builder.AddMwAdapter(authGrop)

	return nil
}
