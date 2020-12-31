// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/chinmobi/gin-mvc/controller"
	mw "github.com/chinmobi/gin-mvc/middleware"

	"github.com/gin-gonic/gin"
)

func setUpRoutes(router *gin.RouterGroup, mws *mw.MiddlewareSet, ctrls *ControllerSet) error {
	commonMWs := mws.CommonHandlersChain()
	authMWs := mws.AuthHandlersChain()

	// Set up routes' handlers
	router.Use(commonMWs...)

	v1 := router.Group("/api/v1", authMWs...)
	{
		setupUserRoutes(v1, mws, ctrls.userCtrl)
	}

	router.GET("/", controller.HandleDefault)

	return nil
}
