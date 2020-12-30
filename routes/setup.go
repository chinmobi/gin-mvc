// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware"
	"github.com/chinmobi/gin-mvc/web/ctx"
)

func SetUp(web *ctx.WebContext, app *app.App) error {
	ctrls, err := setUp(web, app)
	if err != nil {
		return err
	}

	mws := getMiddlewares(web)

	return setUpRoutes(web.RootRouter(), ctrls, mws)
}

func TearDown(web *ctx.WebContext) error {
	return tearDown(web)
}

func getMiddlewares(web *ctx.WebContext) *middleware.MiddlewareSet {
	set := web.GetMiddlewares()
	return set.(*middleware.MiddlewareSet)
}

// Manage controllers (setUp / tearDown)

func setUp(web *ctx.WebContext, app *app.App) (*ControllerSet, error) {
	ctrls := NewControllerSet()
	if err := ctrls.setUp(app); err != nil {
		return nil, err
	}
	web.SetControllers(ctrls)

	return ctrls, nil
}

func tearDown(web *ctx.WebContext) error {
	set := web.GetControllers()
	if set != nil {
		if ctrls, ok := set.(*ControllerSet); ok {
			web.SetControllers(nil)

			if err := ctrls.tearDown(); err != nil {
				return err
			}
		}
	}
	return nil
}
