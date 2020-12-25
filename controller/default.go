// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "TODO",
	})
}
