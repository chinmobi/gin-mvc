// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/controller"
	"github.com/chinmobi/gin-mvc/web/ctx"
)

func SetUp(web *ctx.WebContext, app *app.App) error {
	ctrls := NewControllerSet()
	if err := ctrls.setUp(); err != nil {
		return err
	}
	web.SetControllers(ctrls)

	// Set up routes' handlers

	r := web.RootRouter()
	r.GET("/", controller.HandleDefault)
	return nil
}

func TearDown(web *ctx.WebContext) error {
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
