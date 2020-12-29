// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	"fmt"
)

type ApiErrorEntity struct {
	StatusCode   uint     `json:"statusCode"`
	Code         int      `json:"code"`
	Message      string   `json:"message"`
	Errors       []error  `json:"errors"`
}

func NewApiErrorEntity(statusCode uint, err error) *ApiErrorEntity {
	code := int(statusCode)
	return NewApiErrorEntityS(code, err.Error(), err).SetStatusCode(statusCode)
}

func NewApiErrorEntityA(code int, err error) *ApiErrorEntity {
	return NewApiErrorEntityS(code, err.Error(), err)
}

func NewApiErrorEntityS(code int, msg string, errs ...error) *ApiErrorEntity {
	apiErr := &ApiErrorEntity{
		StatusCode:  500,
		Code:        code,
		Message:     msg,
	}

	apiErr.Errors = make([]error, 0, 1)

	apiErr.AddError(errs...)

	return apiErr
}

func (apiErr *ApiErrorEntity) Error() string {
	return fmt.Sprintf("{statusCode: %d, code: %d, message: %s}", apiErr.StatusCode, apiErr.Code, apiErr.Message)
}

func (apiErr *ApiErrorEntity) AddError(errs ...error) *ApiErrorEntity {
	apiErr.Errors = append(apiErr.Errors, errs...)
	return apiErr
}

func (apiErr *ApiErrorEntity) SetStatusCode(statusCode uint) *ApiErrorEntity {
	apiErr.StatusCode = statusCode
	return apiErr
}

func (apiErr *ApiErrorEntity) GetStatusCode() int {
	return int(apiErr.StatusCode)
}
