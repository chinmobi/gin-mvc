// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// The graceful-shutdown codes is based on the gin-gonic example codes
// (graceful-shutdown by Bo-Yi Wu (appleboy)) that can be found
// at https://github.com/gin-gonic/examples

package ctx

import (
	"github.com/gin-gonic/gin"
)

type WebContext struct {
	engine *gin.Engine
}

func NewWebContext(engine *gin.Engine) *WebContext {
	ctx := &WebContext{
		engine: engine,
	}
	return ctx
}

func (ctx *WebContext) RootRouter() *gin.RouterGroup {
	return &ctx.engine.RouterGroup
}
