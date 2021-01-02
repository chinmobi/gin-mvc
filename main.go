// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/web"
)

func main() {
	app := app.New(config.Default())

	web.StartServer(app)
}
