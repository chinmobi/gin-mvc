// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/controller"
	"github.com/chinmobi/gin-mvc/security"
)

type ControllerSet struct {
	userCtrl  *controller.UserController
}

func NewControllerSet() *ControllerSet {
	set := &ControllerSet{
	}
	return set
}

func (set *ControllerSet) setUp(app *app.App, authHandler security.AuthHandler) error {
	set.userCtrl = controller.NewUserController(app.ServiceSupplier())
	return nil
}

func (set *ControllerSet) tearDown() error {
	return nil
}
