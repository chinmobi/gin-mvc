// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !fasthttp

package srv

import (
	"context"
	"net/http"
)

type Server struct {
	httpSrv http.Server
}

func NewServer(config *ServerConfig) *Server {
	server := &Server{
		httpSrv: http.Server{
			Addr: config.Addr,
			Handler: config.Handler,
		},
	}
	return server
}

func (srv *Server) Close() error {
	return srv.httpSrv.Close()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpSrv.Shutdown(ctx)
}

func (srv *Server) ListenAndServe() error {
	return srv.httpSrv.ListenAndServe()
}
