// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"github.com/chinmobi/gin-mvc/security"
)

// Context holder used for per request, NOT goroutine safe!

const CTX_SECURITY_HOLDER = "CTX_SECURITY_HOLDER"

type SecurityContextHolder interface {
	GetSecurityContex() security.SecurityContext
}

type ContextHolder struct {
	context   Context
	security  security.Context
}

func NewContextHolder() *ContextHolder {
	h := &ContextHolder{
	}
	return h
}

func (h *ContextHolder) GetSecurityContex() security.SecurityContext {
	return &h.security
}

func (h *ContextHolder) GetSecurity() *security.Context {
	return &h.security
}

func (h *ContextHolder) GetSecurityContextHolder() SecurityContextHolder {
	return h
}

func (h *ContextHolder) GetContext() *Context {
	return &h.context
}
