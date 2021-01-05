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
	"github.com/chinmobi/gin-mvc/web/srv"

	"github.com/gin-gonic/gin"
)

type ServerContext struct {
	app     *app.App
	engine  *gin.Engine
	web     *ctx.WebContext
}

func NewServerContext(app *app.App) *ServerContext {
	// Starts a new Gin instance with no middle-ware
	engine := gin.New()

	server := &ServerContext{
		app: app,
		engine: engine,
		web: ctx.NewWebContext(engine),
	}
	return server
}

func (s *ServerContext) SetUp() {
	if err := s.app.Start(); err != nil {
		log.Fatalf("App starting fault: %+v\n", err)
	}

	// Set up middlewares
	if err := middleware.SetUp(s.web, s.app); err != nil {
		shutDownApp(s.app)
		log.Fatalf("Setup web middlewares fault: %+v", err)
	}

	// Define handlers
	if err := routes.SetUp(s.web, s.app); err != nil {
		tearDownMiddleware(s.web)
		shutDownApp(s.app)
		log.Fatalf("Setup web routers fault: %+v", err)
	}

	log.Println("Server started")
}

func (s *ServerContext) TearDown() {
	log.Println("Server exiting...")

	tearDownWebContext(s.web)
	shutDownApp(s.app)

	log.Println("Server exited")
}

func (s *ServerContext) Engine() *gin.Engine {
	return s.engine
}

func (s *ServerContext) App() *app.App {
	return s.app
}

// Start the web server

func StartServer(app *app.App) {

	server := NewServerContext(app)
	server.SetUp()

	web := server.web

	// Listen and serve on defined port
	srv := srv.NewServer(&srv.ServerConfig{
		Addr:    ":" + app.Config().Server.Port,
		Handler: server.Engine(),
	})

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			tearDownWebContext(web)
			shutDownApp(app)
			log.Fatalf("Http server listen fault: %+v", err)
		}
	}()

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
