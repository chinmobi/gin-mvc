// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

	"github.com/chinmobi/ginmod/restful"
	"github.com/chinmobi/gin-mvc/security/web/auth"

	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "v1"
)

func RespAuthenticationError(c *gin.Context, err *auth.ErrAuthentication) (bool, error) {
	apiErr := restful.NewApiErrorEntity(http.StatusUnauthorized, err)
	c.JSON(apiErr.GetStatusCode(), restful.CreateApiErrorBody(apiVersion, apiErr))
	return true, nil
}
