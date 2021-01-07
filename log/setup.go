// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package log

import (
	"github.com/chinmobi/gin-mvc/config"
)

func SetUp(config *config.Logger) error {
	return setUpZap(config)
}
