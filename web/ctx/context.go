// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"github.com/gin-gonic/gin"
)

type WebContext struct {
	engine       *gin.Engine
	controllers  interface{} // avoid importing cycle
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

func (ctx *WebContext) SetControllers(controllers interface{}) {
	ctx.controllers = controllers
}

func (ctx *WebContext) GetControllers() interface{} {
	return ctx.controllers
}
