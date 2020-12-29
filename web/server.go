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

	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/middleware"
	"github.com/chinmobi/gin-mvc/routes"
	"github.com/chinmobi/gin-mvc/web/ctx"

	"github.com/gin-gonic/gin"
)

func StartServer(app *app.App) {

	// Starts a new Gin instance with no middle-ware
	engine := gin.New()

	web := ctx.NewWebContext(engine)

	// Set up middlewares
	if err := middleware.SetUp(web, app); err != nil {
		shutDownApp(app)
		log.Fatalf("Setup web middlewares fault: %+v", err)
	}

	// Define handlers
	if err := routes.SetUp(web, app); err != nil {
		tearDownMiddleware(web)
		shutDownApp(app)
		log.Fatalf("Setup web routers fault: %+v", err)
	}

	// Listen and serve on defined port
	srv := &http.Server{
		Addr:    ":" + app.Config().Server.Port,
		Handler: engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			tearDownWebContext(web)
			shutDownApp(app)
			log.Fatalf("Http server listen fault: %+v", err)
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
		tearDownWebContext(web)
		shutDownApp(app)
		log.Println("Server exited")
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %+v", err)
	}

	log.Println("Server exiting...")
}

func tearDownWebContext(web *ctx.WebContext) {
	if err := routes.TearDown(web); err != nil {
		log.Printf("Web routes tearing down: %+v", err)
	}
	tearDownMiddleware(web)
}

func tearDownMiddleware(web *ctx.WebContext) {
	if err := middleware.TearDown(web); err != nil {
		log.Printf("Web middlewares tearing down: %+v", err)
	}
}

func shutDownApp(app *app.App) {
	if err := app.Shutdown(); err != nil {
		log.Printf("App shutting down: %+v", err)
	}
}
