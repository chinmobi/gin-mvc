// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"errors"

	"github.com/chinmobi/gin-mvc/ctx"

	"github.com/gin-gonic/gin"
)

type SecurityInterceptor struct {
	permissions     *PermissionsGroup
	deniedHandler   AccessDeniedHandler
	evaluator       PrivilegeEvaluator
}

type InterceptorAgent struct {
	interceptor     *SecurityInterceptor
	controllerFunc  gin.HandlerFunc
}

func (si *SecurityInterceptor) InterceptAgent(controllerFunc gin.HandlerFunc) gin.HandlerFunc {
	agent := &InterceptorAgent{
		interceptor: si,
		controllerFunc: controllerFunc,
	}
	return agent.getHandlerFunc()
}

func (agent *InterceptorAgent) getHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		agent.doIntercept(c)
	}
}

func (agent *InterceptorAgent) doIntercept(c *gin.Context) {
	if err := agent.interceptor.decide(c); err != nil {
		agent.interceptor.deniedHandler.OnAccessDenied(c, NewErrAccessDenied(err))
		return
	}
	agent.controllerFunc(c)
}

func (si *SecurityInterceptor) decide(c *gin.Context) error {
	permissions := si.permissions
	if permissions == nil {
		ctxPerms, exists := c.Get(CTX_ACCESS_PERMISSIONS)
		if !exists {
			return nil
		}

		perms, ok := ctxPerms.(*PermissionsGroup)
		if !ok {
			return errors.New("Invalid permissions")
		}
		permissions = perms
	}

	holder, exists := c.Get(ctx.CTX_SECURITY_HOLDER)
	if !exists {
		return errors.New("SecurityContextHolder not exists")
	}

	securityHolder, ok := holder.(ctx.SecurityContextHolder)
	if !ok {
		return errors.New("Invalid SecurityContextHolder")
	}

	authentication := securityHolder.GetSecurityContex().GetAuthentication()
	if authentication == nil {
		return errors.New("Nil authentication")
	}

	if !authentication.IsAuthenticated() {
		return errors.New("Not authenticated")
	}

	if !si.evaluator.IsAllowed(authentication, permissions) {
		return errors.New("Not allowed permissions")
	}

	return nil
}

// InterceptorBuilder

type InterceptorBuilder struct {
	permissions   *PermissionsGroup
	deniedHandler AccessDeniedHandler
	evaluator     PrivilegeEvaluator
}

func NewBuilder(onAccessDenied OnAccessDeniedFunc) *InterceptorBuilder {
	builder := &InterceptorBuilder{
		deniedHandler: WrapAccessDeniedFunc(onAccessDenied),
		evaluator: SimplePrivilegeEvaluator{},
	}

	return builder
}

func (b *InterceptorBuilder) SetPermissions(permissions *PermissionsGroup) {
	b.permissions = permissions
}

func (b *InterceptorBuilder) SetDeniedHandler(handler AccessDeniedHandler) {
	b.deniedHandler = handler
}

func (b *InterceptorBuilder) SetEvaluator(evaluator PrivilegeEvaluator) {
	b.evaluator = evaluator
}

func (b *InterceptorBuilder) Build() *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    b.permissions,
		deniedHandler:  b.deniedHandler,
		evaluator:      b.evaluator,
	}
	return si
}
