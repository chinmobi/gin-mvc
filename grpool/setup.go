// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpool

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/grpool/gr"
)

func SetUp(config *config.Grpool) (gr.ExecutorService, error) {
	return newAntsExecutor(config)
}
