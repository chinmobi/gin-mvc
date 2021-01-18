// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/modlib/evt/event"
)

type Supplier interface {
	GetEventBroker() event.Broker
	GetAuthService() *AuthService
	GetUserService() *UserService
}

func SetUp(models model.Supplier, broker event.Broker) (Supplier, error) {
	return createSupplier(models, broker)
}

func TearDown(services Supplier) error {
	if ss, ok := services.(*serviceSupplier); ok {
		return ss.release()
	}
	return nil
}
