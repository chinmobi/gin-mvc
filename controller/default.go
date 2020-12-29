// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"github.com/chinmobi/gin-mvc/restful"

	"github.com/gin-gonic/gin"
)

func HandleDefault(c *gin.Context) {
	restful.RespDataEntity(c, gin.H{
		"message": "TODO",
	})
}
