// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security/auth"
	"github.com/chinmobi/gin-mvc/security/ctx"

	"github.com/gin-gonic/gin"
)

type AuthenticationHelper interface {
	AttemptAuthentication(c *gin.Context) (Authentication, error)
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

func (ap *AuthProcessor) AuthManager() auth.AuthManager {
	return ap.base.manager
}

func (ap *AuthProcessor) AuthHandler() AuthHandler {
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

		done, _ := ap.AuthHandler().OnAuthFailure(c, authErr)
		return !done
	}

	return true
}

func (ap *AuthProcessor) doProcess(c *gin.Context, s ctx.SecurityContext) error {
	auth := s.GetAuthentication()
	if auth != nil && auth.IsAuthenticated() {
		return nil
	}

	auth, err := ap.helper.AttemptAuthentication(c)
	if err != nil {
		return err
	}
	if auth == nil {
		return nil
	}

	s.SetAuthentication(auth)

	if auth.IsAuthenticated() {
		return nil
	}

	result, err := ap.AuthManager().Authenticate(auth)
	if err != nil {
		return err
	}

	if result != nil {
		s.SetAuthentication(result)
	}

	return nil
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
