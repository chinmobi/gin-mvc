// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build fasthttp

package srv

import (
	"context"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type Server struct {
	httpSrv fasthttp.Server
	Addr string
}

func NewServer(config *ServerConfig) *Server {
	server := &Server{
		httpSrv: fasthttp.Server{
			Handler: fasthttpadaptor.NewFastHTTPHandler(config.Handler),
			Name: "gin-mvc",
		},
		Addr: config.Addr,
	}
	return server
}

func (srv *Server) Close() error {
	return srv.httpSrv.Shutdown()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	// FIX ME: close the KeepAlived connections before invoking the shutdown!
	return srv.httpSrv.Shutdown()
}

func (srv *Server) ListenAndServe() error {
	return srv.httpSrv.ListenAndServe(srv.Addr)
}
