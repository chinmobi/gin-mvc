// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"errors"

	"github.com/chinmobi/gin-mvc/security/consts"

	"github.com/gin-gonic/gin"
)

func SetSecurityHolder(c *gin.Context) SecurityContextHolder {
	holder, exists := c.Get(CTX_SECURITY_HOLDER)
	if exists && holder != nil {
		if securityHolder, ok := holder.(SecurityContextHolder); ok {
			return securityHolder
		}
	}

	securityHolder := NewContextHolder().GetSecurityContextHolder()

	c.Set(CTX_SECURITY_HOLDER, securityHolder)

	return securityHolder
}

func GetSecurityHolder(c *gin.Context) (SecurityContextHolder, error) {
	holder, exists := c.Get(CTX_SECURITY_HOLDER)
	if !exists || holder == nil {
		return nil, errors.New(consts.ERR_STR_HOLDER_NOT_EXISTS)
	}

	securityHolder, ok := holder.(SecurityContextHolder)
	if !ok {
		return nil, errors.New(consts.ERR_STR_INVALID_HOLDER)
	}

	return securityHolder, nil
}

func GetSecurityContext(c *gin.Context) (SecurityContext, error) {
	securityHolder, err := GetSecurityHolder(c)
	if err != nil {
		return nil, err
	}

	return securityHolder.GetSecurityContex(), nil
}

func CleanSecurityHolder(c *gin.Context) {
	c.Set(CTX_SECURITY_HOLDER, nil)
}
