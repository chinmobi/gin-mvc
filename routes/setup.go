// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/controller"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup, app *app.App) {
	r.GET("/", controller.HandleDefault)
}
