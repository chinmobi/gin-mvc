// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/ctx"
	sec "github.com/chinmobi/gin-mvc/security"

	"github.com/gin-gonic/gin"
)

type AuthenticationHelper interface {
	AttemptAuthentication(c *gin.Context, handler sec.AuthSuccessHandler) (sec.Authentication, error)
	TearDown() error
}

type AuthHelper = AuthenticationHelper

type AuthenticationProcessor struct {
	base        *AuthProcessorBase
	helper       AuthHelper
	handlerFunc  gin.HandlerFunc
}

type AuthProcessor = AuthenticationProcessor

func NewAuthProcessor(base *AuthProcessorBase, helper AuthHelper) *AuthProcessor {
	processor := &AuthProcessor{
		base: base,
		helper: helper,
	}
	return processor
}

func (ap *AuthProcessor) AuthManager() AuthManager {
	return ap.base.manager
}

func (ap *AuthProcessor) AuthHandler() sec.AuthHandler {
	return ap.base.authHandler
}

func (ap *AuthProcessor) process(c *gin.Context) {
	if ap.processAuth(c) {
		c.Next()
	}
}

func (ap *AuthProcessor) processAuth(c *gin.Context) bool {
	securityContext := ctx.SetSecurityHolder(c).GetSecurityContex()

	err := ap.doProcess(c, securityContext)
	if err != nil {
		securityContext.CleanAuthentication()

		authErr, ok := err.(*ErrAuthentication)
		if !ok {
			authErr = NewErrAuthentication(err)
		}

		ap.AuthHandler().OnAuthFailure(c, authErr)
		return false
	}

	return true
}

func (ap *AuthProcessor) doProcess(c *gin.Context, s sec.SecurityContext) error {
	auth := s.GetAuthentication()
	if auth != nil && auth.IsAuthenticated() {
		return nil
	}

	auth, err := ap.helper.AttemptAuthentication(c, ap.AuthHandler())
	if err != nil {
		return err
	}
	if auth == nil {
		return nil
	}

	if auth.IsAuthenticated() {
		// The AuthSuccessHandler SHOULD be invoked within the helper's AttemptAuthentication method!
		s.SetAuthentication(auth)
		return nil
	}

	result, err := ap.AuthManager().Authenticate(auth)
	if err != nil {
		return err
	}

	if result == nil {
		if auth.IsAuthenticated() {
			ap.AuthHandler().OnAuthSuccess(c, auth)
		}
		return nil
	}

	if !result.IsAuthenticated() {
		s.SetAuthentication(result)
	} else {
		ap.successfulAuth(c, s, result)
	}

	return nil
}

func (ap *AuthProcessor) successfulAuth(c *gin.Context, s sec.SecurityContext, auth sec.Authentication) {
	s.SetAuthentication(auth)

	ap.AuthHandler().OnAuthSuccess(c, auth)
}

func (ap *AuthProcessor) HandlerFunc() gin.HandlerFunc {
	if ap.handlerFunc == nil {
		ap.handlerFunc = ap.createHandlerFunc()
	}
	return ap.handlerFunc
}

func (ap *AuthProcessor) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ap.process(c)
	}
}

func (ap *AuthProcessor) TearDown() error {
	return ap.helper.TearDown()
}
