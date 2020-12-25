// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	port := "8080"

	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define handlers
	r.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TODO",
		})
	})

	// Listen and serve on defined port
	r.Run(":" + port)
}
