// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"errors"

	"github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/consts"

	"github.com/gin-gonic/gin"
)

func GetSecurityHolder(c *gin.Context) (SecurityContextHolder, error) {
	holder, exists := c.Get(CTX_SECURITY_HOLDER)
	if !exists {
		return nil, errors.New(consts.ERR_STR_HOLDER_NOT_EXISTS)
	}

	securityHolder, ok := holder.(SecurityContextHolder)
	if !ok {
		return nil, errors.New(consts.ERR_STR_INVALID_HOLDER)
	}

	return securityHolder, nil
}

func GetSecurityContext(c *gin.Context) (security.SecurityContext, error) {
	securityHolder, err := GetSecurityHolder(c)
	if err != nil {
		return nil, err
	}

	return securityHolder.GetSecurityContex(), nil
}
