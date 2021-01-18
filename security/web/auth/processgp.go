// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"errors"

	myerrors "github.com/chinmobi/ginmod/errors"
	"github.com/chinmobi/gin-mvc/security/consts"
	"github.com/chinmobi/gin-mvc/security/ctx"

	"github.com/gin-gonic/gin"
)

type ProcessorGroup struct {
	configurer   *ProcessorConfigurer
	processors   []*AuthProcessor
	handlerFunc  gin.HandlerFunc
}

func NewProcessorGroup(config *ProcessorConfigurer) *ProcessorGroup {
	group := &ProcessorGroup{
		configurer: config,
	}
	return group
}

func (pg *ProcessorGroup) Configurer() *ProcessorConfigurer {
	return pg.configurer
}

func (pg *ProcessorGroup) CreateProcessor(helper AuthHelper) *AuthProcessor {
	base := pg.configurer.ProcessorBase()
	return NewAuthProcessor(base, helper)
}

func (pg *ProcessorGroup) AddProcessor(processor ...*AuthProcessor) {
	pg.processors = append(pg.processors, processor...)
}

func (pg *ProcessorGroup) process(c *gin.Context) {
	if pg.processAuth(c) {
		c.Next()
	}
}

func (pg *ProcessorGroup) processAuth(c *gin.Context) bool {
	err := pg.processGroup(c)

	if err != nil {
		authErr, ok := err.(*ErrAuthentication)
		if !ok {
			authErr = NewErrAuthentication(err)
		}

		done, _ := pg.authHandler().OnAuthFailure(c, authErr)
		return !done
	}

	return true
}

func (pg *ProcessorGroup) processGroup(c *gin.Context) error {
	securityContext := ctx.SetSecurityHolder(c).GetSecurityContex()

	if err := pg.doProcess(c, securityContext); err != nil {
		securityContext.CleanAuthentication()
		return err
	}

	auth := securityContext.GetAuthentication()
	if auth != nil && auth.IsAuthenticated() {
		return nil
	}

	securityContext.CleanAuthentication()
	return errors.New(consts.ERR_STR_AUTHENTICATION_FAILED)
}

func (pg *ProcessorGroup) doProcess(c *gin.Context, s ctx.SecurityContext) error {
	for i, cnt := 0, len(pg.processors); i < cnt; i++ {
		processor := pg.processors[i]

		if err := processor.doProcess(c, s); err != nil {
			return err
		}

		auth := s.GetAuthentication()
		if auth != nil && auth.IsAuthenticated() {
			break
		}
	}
	return nil
}

func (pg *ProcessorGroup) authHandler() AuthHandler {
	return pg.configurer.AuthHandlerSet()
}

func (pg *ProcessorGroup) HandlerFunc() gin.HandlerFunc {
	if pg.handlerFunc == nil {
		pg.handlerFunc = pg.createHandlerFunc()
	}
	return pg.handlerFunc
}

func (pg *ProcessorGroup) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		pg.process(c)
	}
}

func (pg *ProcessorGroup) TearDown() error {
	errs := myerrors.NewWrapErrorsErr()

	for i := len(pg.processors)-1; i >= 0; i-- {
		processor := pg.processors[i]
		if err := processor.TearDown(); err != nil {
			errs.Wrap(err)
		}
	}

	return errs.ToError()
}
