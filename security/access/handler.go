// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/security/errors"

	"github.com/gin-gonic/gin"
)

type ErrAccessDenied = errors.ErrAccessDenied

type OnAccessDeniedFunc func(c *gin.Context, err *ErrAccessDenied) error

type AccessDeniedHandler interface {
	OnAccessDenied(c *gin.Context, err *ErrAccessDenied) (bool, error)
}

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

func NewErrAccessDenied(cause error) *ErrAccessDenied {
	return errors.NewErrAccessDenied(cause)
}
