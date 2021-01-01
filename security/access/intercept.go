// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"errors"

	"github.com/chinmobi/gin-mvc/ctx"
	"github.com/chinmobi/gin-mvc/security/consts"

	"github.com/gin-gonic/gin"
)

type SecurityInterceptor struct {
	permissions    *PermissionsGroup
	deniedHandler  AccessDeniedHandler
	evaluator      PrivilegeEvaluator
	decisionAgent  *AccessDecisionAgent
}

type AccessDecisionAgent struct {
	interceptor    *SecurityInterceptor
	handlerFunc    gin.HandlerFunc
}

func (si *SecurityInterceptor) DecisionAgent() *AccessDecisionAgent {
	if si.decisionAgent != nil {
		return si.decisionAgent
	}

	agent := &AccessDecisionAgent{
		interceptor: si,
	}
	si.decisionAgent = agent

	return agent
}

func (agent *AccessDecisionAgent) DecideHandlerFunc() gin.HandlerFunc {
	if agent.handlerFunc == nil {
		agent.handlerFunc = agent.createHandlerFunc()
	}
	return agent.handlerFunc
}

func (agent *AccessDecisionAgent) DecideControllerFunc(ctrlFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if agent.doDecide(c) {
			ctrlFunc(c)
		}
	}
}

func (agent *AccessDecisionAgent) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if agent.doDecide(c) {
			c.Next()
		}
	}
}

func (agent *AccessDecisionAgent) doDecide(c *gin.Context) bool {
	if err := agent.interceptor.decide(c); err != nil {
		agent.interceptor.deniedHandler.OnAccessDenied(c, NewErrAccessDenied(err))
		return false
	}
	return true
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
			return errors.New(consts.ERR_STR_INVALID_PERMISSIONS)
		}
		permissions = perms
	}

	holder, exists := c.Get(ctx.CTX_SECURITY_HOLDER)
	if !exists {
		return errors.New(consts.ERR_STR_HOLDER_NOT_EXISTS)
	}

	securityHolder, ok := holder.(ctx.SecurityContextHolder)
	if !ok {
		return errors.New(consts.ERR_STR_INVALID_HOLDER)
	}

	authentication := securityHolder.GetSecurityContex().GetAuthentication()
	if authentication == nil {
		return errors.New(consts.ERR_STR_NIL_AUTHENTICATION)
	}

	if !authentication.IsAuthenticated() {
		return errors.New(consts.ERR_STR_NOT_AUTHENTICATED)
	}

	if !si.evaluator.IsAllowed(authentication, permissions) {
		return errors.New(consts.ERR_STR_NOT_ALLOWED_PERMISSIONS)
	}

	return nil
}
