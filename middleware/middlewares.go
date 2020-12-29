// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/chinmobi/gin-mvc/app"
)

type MiddlewareSet struct {
}

func NewMiddlewareSet() *MiddlewareSet {
	set := &MiddlewareSet{
	}
	return set
}

func (set *MiddlewareSet) setUp(app *app.App) error {
	return nil
}

func (set *MiddlewareSet) tearDown() error {
	return nil
}
