// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/security/errors"

	"github.com/gin-gonic/gin"
)

type ErrAccessDenied = errors.ErrAccessDenied

// AccessDeniedHandler

type OnAccessDeniedFunc func(c *gin.Context, err *ErrAccessDenied) error

type AccessDeniedHandler interface {
	OnAccessDenied(c *gin.Context, err *ErrAccessDenied) (bool, error)
}

// accessDeniedFuncWrap

type accessDeniedFuncWrap struct {
	onAccessDenied OnAccessDeniedFunc
}

func (w *accessDeniedFuncWrap) OnAccessDenied(c *gin.Context, err *ErrAccessDenied) (bool, error) {
	if w.onAccessDenied != nil {
		return true, w.onAccessDenied(c, err)
	}
	return false, nil
}

func WrapAccessDeniedFunc(onAccessDenied OnAccessDeniedFunc) AccessDeniedHandler {
	wrap := &accessDeniedFuncWrap{
		onAccessDenied: onAccessDenied,
	}
	return wrap
}

// NullAccessDeniedHandler

type NullAccessDeniedHandler struct{}

func (n NullAccessDeniedHandler) OnAccessDenied(c *gin.Context, err *ErrAccessDenied) (bool, error) {
	return false, nil
}

// Wrap the errors.NewErrAccessDenied

func NewErrAccessDenied(cause error) *ErrAccessDenied {
	return errors.NewErrAccessDenied(cause)
}
