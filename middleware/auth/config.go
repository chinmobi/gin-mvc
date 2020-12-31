// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security"
)

func Configure(builder *mw.Builder, authHandler security.AuthHandlerSetter, app *app.App) error {
	setUp(authHandler)

	// The CtxHolder MUST be the first handler of all the security auth handlers
	builder.AddMwAdapter(NewCtxHolder())
	return nil
}
