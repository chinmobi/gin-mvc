// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/web"
)

const (
	STARTING_FAULT_LOG_FMT = "App starting fault: %+v\n"
)

func main() {
	app := app.New(config.Default())

	if err := app.Start(); err != nil {
		log.Fatalf(STARTING_FAULT_LOG_FMT, err)
	}

	web.StartServer(app)
}
