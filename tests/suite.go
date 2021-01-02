// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tests

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/web"

	"github.com/gin-gonic/gin"
)

var serverContext *web.ServerContext

func SetUp() *gin.Engine {
	if serverContext == nil {
		app := app.New(config.Default())

		server := web.NewServerContext(app)
		server.SetUp()

		serverContext = server
	}

	return serverContext.Engine()
}

func TearDown() {
	if serverContext != nil {
		server := serverContext
		serverContext = nil

		server.TearDown()
	}
}

func GinEngine() *gin.Engine {
	if serverContext != nil {
		return serverContext.Engine()
	}
	return nil
}
