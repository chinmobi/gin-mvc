// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

	"github.com/chinmobi/gin-mvc/restful"
	"github.com/chinmobi/gin-mvc/security"

	"github.com/gin-gonic/gin"
)

func RespAuthenticationError(c *gin.Context, err *security.ErrAuthentication) error {
	apiErr := restful.NewApiErrorEntity(http.StatusUnauthorized, err)
	c.JSON(apiErr.GetStatusCode(), restful.CreateApiErrorBody(apiErr))
	return nil
}
