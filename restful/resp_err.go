// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	"net/http"

	"github.com/chinmobi/gin-mvc/errors"

	"github.com/gin-gonic/gin"
)

func RespBadRequest(c *gin.Context, err error) {
	err = errors.NewErrBadRequest(err)

	apiErr := NewApiErrorEntity(http.StatusBadRequest, err)
	c.JSON(apiErr.GetStatusCode(), CreateApiErrorBody(apiErr))
}

func RespServiceError(c *gin.Context, err error) {
	var statusCode uint = http.StatusInternalServerError

	switch err.(type) {
		case *errors.ErrNotFound:
			statusCode = http.StatusNotFound

		case *errors.ErrAlreadyExists, *errors.ErrBadRequest:
			statusCode = http.StatusBadRequest

		case *errors.ErrLackOfParameter, *errors.ErrInvalidParameter:
			statusCode = http.StatusBadRequest

		case *errors.ErrMethodNotAllowed:
			statusCode = http.StatusMethodNotAllowed

		case *errors.ErrNotImplemented:
			statusCode = http.StatusNotImplemented

		case *errors.ErrInternalError:
			// Nothing to do.

		default:
			err = errors.NewErrInternalError(err)
	}

	apiErr := NewApiErrorEntity(statusCode, err)
	c.JSON(apiErr.GetStatusCode(), CreateApiErrorBody(apiErr))
}
