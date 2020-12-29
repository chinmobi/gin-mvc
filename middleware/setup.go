// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/web/ctx"
)

func SetUp(web *ctx.WebContext, app *app.App) error {
	mws := NewMiddlewareSet()
	if err := mws.setUp(app); err != nil {
		return err
	}
	web.SetMiddlewares(mws)

	return nil
}

func TearDown(web *ctx.WebContext) error {
	set := web.GetMiddlewares()
	if set != nil {
		if mws, ok := set.(*MiddlewareSet); ok {
			web.SetMiddlewares(nil)

			if err := mws.tearDown(); err != nil {
				return err
			}
		}
	}
	return nil
}
