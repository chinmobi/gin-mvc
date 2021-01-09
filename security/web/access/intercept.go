// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"errors"

	"github.com/chinmobi/gin-mvc/security/access"
	"github.com/chinmobi/gin-mvc/security/consts"
	"github.com/chinmobi/gin-mvc/security/ctx"

	"github.com/gin-gonic/gin"
)

const CTX_ACCESS_PERMISSIONS = "CTX_ACCESS_PERMISSIONS"

type SecurityInterceptor struct {
	permissions    *access.PermissionsGroup
	deniedHandler  AccessDeniedHandler
	evaluator      access.PrivilegeEvaluator
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

		perms, ok := ctxPerms.(*access.PermissionsGroup)
		if !ok {
			return errors.New(consts.ERR_STR_INVALID_PERMISSIONS)
		}
		permissions = perms
	}

	securityHolder, err := ctx.GetSecurityHolder(c)
	if err != nil {
		return err
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
