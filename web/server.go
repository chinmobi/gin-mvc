// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// The graceful-shutdown codes is based on the gin-gonic example codes
// (graceful-shutdown by Bo-Yi Wu (appleboy)) that can be found
// at https://github.com/gin-gonic/examples

package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chinmobi/gin-mvc/config"

	"github.com/gin-gonic/gin"
)

func StartServer(config *config.Config) {
	port := config.Server.Port

	// Starts a new Gin instance with no middle-ware
	router := gin.New()

	// Define handlers
	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "TODO",
		})
	})

	// Listen and serve on defined port
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
		log.Println("Server exited")
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: %+v", err)
	}

	log.Println("Server exiting...")
}
