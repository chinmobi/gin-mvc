// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/errors"

	"github.com/gin-gonic/gin"
)

type Authentication = security.Authentication

type ErrAuthentication = errors.ErrAuthentication

type OnAuthSuccessFunc func(c *gin.Context, auth Authentication) (bool, error)
type OnAuthFailureFunc func(c *gin.Context, authErr *ErrAuthentication) (bool, error)

type AuthSuccessHandler interface {
	OnAuthSuccess(c *gin.Context, auth Authentication) (bool, error)
}

type AuthFailureHandler interface {
	OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (bool, error)
}

type AuthHandler interface {
	AuthSuccessHandler
	AuthFailureHandler
}

type AuthHandlerSetter interface {
	AddAuthSuccessFunc(onSuccess ...OnAuthSuccessFunc)
	AddAuthFailureFunc(onFailure ...OnAuthFailureFunc)
}

type AuthHandlerSet struct {
	successFuncChain  []OnAuthSuccessFunc
	failureFuncChain  []OnAuthFailureFunc
}

func NewAuthHandlerSet() *AuthHandlerSet {
	set := &AuthHandlerSet{
	}
	return set
}

// AuthHandlerSetter methods

func (set *AuthHandlerSet) AddAuthSuccessFunc(onSuccess ...OnAuthSuccessFunc) {
	set.successFuncChain = append(set.successFuncChain, onSuccess...)
}

func (set *AuthHandlerSet) AddAuthFailureFunc(onFailure ...OnAuthFailureFunc) {
	set.failureFuncChain = append(set.failureFuncChain, onFailure...)
}

func (set *AuthHandlerSet) Clear() {
	if set.successFuncChain != nil {
		set.successFuncChain = set.successFuncChain[0:0]
	}
	if set.failureFuncChain != nil {
		set.failureFuncChain = set.failureFuncChain[0:0]
	}
}

// AuthHandler methods

func (set *AuthHandlerSet) OnAuthSuccess(c *gin.Context, auth Authentication) (bool, error) {
	for i := len(set.successFuncChain)-1; i >= 0; i-- {
		onSuccess := set.successFuncChain[i]

		done, err := onSuccess(c, auth)
		if err != nil {
			return done, err
		}
		if done {
			return true, nil
		}
	}
	return false, nil
}

func (set *AuthHandlerSet) OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (bool, error) {
	for i := len(set.failureFuncChain)-1; i >= 0; i-- {
		onFailure := set.failureFuncChain[i]

		done, err := onFailure(c, authErr)
		if err != nil {
			return done, err
		}
		if done {
			return true, nil
		}
	}
	return false, nil
}

// Wrap the errors.NewErrAuthentication

func NewErrAuthentication(cause error) *ErrAuthentication {
	return errors.NewErrAuthentication(cause)
}
