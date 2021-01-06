// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package evt

import (
	"github.com/chinmobi/gin-mvc/evt/event"
	"github.com/chinmobi/gin-mvc/grpool/gr"
)

func SetUp(executor gr.Executor) (event.Broker, error) {
	return newEngine(newMulticaster(executor)), nil
}
