// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/chinmobi/ginmod/restful"

	"github.com/gin-gonic/gin"
)

const (
	defaultApiVersion = "v1"
)

func HandleDefault(c *gin.Context) {
	restful.RespDataEntity(c, defaultApiVersion, gin.H{
		"message": "TODO",
	})
}
