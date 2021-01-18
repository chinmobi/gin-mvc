// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"net/http"

	"github.com/chinmobi/ginmod/restful"
	"github.com/chinmobi/gin-mvc/security/web/access"

	"github.com/gin-gonic/gin"
)

const (
	apiVersion = "v1"
)

func RespAccessDenied(c *gin.Context, err *access.ErrAccessDenied) error {
	apiErr := restful.NewApiErrorEntity(http.StatusForbidden, err)
	c.JSON(apiErr.GetStatusCode(), restful.CreateApiErrorBody(apiVersion, apiErr))
	return nil
}
