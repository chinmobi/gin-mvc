// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"github.com/gin-gonic/gin"
)

type OnAuthSuccessFunc func(c *gin.Context, auth Authentication) error
type OnAuthFailureFunc func(c *gin.Context, authErr *ErrAuthentication) error

type AuthHandler interface {
	OnAuthSuccess(c *gin.Context, auth Authentication) (int, error)
	OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (int, error)
}

type AuthHandlerSetter interface {
	AddSuccessFunc(onSuccess ...OnAuthSuccessFunc)
	AddFailureFunc(onFailure ...OnAuthFailureFunc)
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

func (set *AuthHandlerSet) AddSuccessFunc(onSuccess ...OnAuthSuccessFunc) {
	set.successFuncChain = append(set.successFuncChain, onSuccess...)
}

func (set *AuthHandlerSet) AddFailureFunc(onFailure ...OnAuthFailureFunc) {
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

func (set *AuthHandlerSet) OnAuthSuccess(c *gin.Context, auth Authentication) (int, error) {
	count := 0
	for i := len(set.successFuncChain)-1; i >= 0; i-- {
		onSuccess := set.successFuncChain[i]

		if err := onSuccess(c, auth); err != nil {
			return count, err
		}

		count++
	}
	return count, nil
}

func (set *AuthHandlerSet) OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (int, error) {
	count := 0
	for i := len(set.failureFuncChain)-1; i >= 0; i-- {
		onFailure := set.failureFuncChain[i]

		if err := onFailure(c, authErr); err != nil {
			return count, err
		}

		count++
	}
	return count, nil
}
