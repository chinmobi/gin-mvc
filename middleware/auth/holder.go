// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/ctx"

	"github.com/gin-gonic/gin"
)

type CtxHolder struct {
}

func NewCtxHolder() *CtxHolder {
	h := &CtxHolder{
	}
	return h
}

func (h *CtxHolder) HandlerFunc() gin.HandlerFunc {
	return handlerFunc()
}

func handlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {

		holder := ctx.NewContextHolder()

		c.Set(ctx.CTX_SECURITY_HOLDER, holder.GetSecurityContextHolder())

		c.Next()

		c.Set(ctx.CTX_SECURITY_HOLDER, nil)
	}
}

func (h *CtxHolder) TearDown() error {
	// Nothing to do
	return nil
}
