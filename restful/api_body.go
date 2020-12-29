// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

const API_VERSION = "v1"


type ApiDataBody struct {
	ApiVersion  string       `json:"apiVersion"`
	Data        interface{}  `json:"data"`
}

func CreateApiDataBody(data interface{}) *ApiDataBody {
	dataBody := &ApiDataBody{
		ApiVersion: API_VERSION,
		Data: data,
	}
	return dataBody
}


type ApiErrorBody struct {
	ApiVersion  string           `json:"apiVersion"`
	Error       *ApiErrorEntity  `json:"error"`
}

func CreateApiErrorBody(err *ApiErrorEntity) *ApiErrorBody {
	errBody := &ApiErrorBody{
		ApiVersion: API_VERSION,
		Error: err,
	}
	return errBody
}
