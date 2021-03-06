// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/modlib/evt/event"
)

type serviceSupplier struct {
	models       model.Supplier
	eventBroker  event.Broker
	authService  *AuthService
	userService  *UserService
}

func createSupplier(models model.Supplier, broker event.Broker) (*serviceSupplier, error) {
	ss := &serviceSupplier{
		models: models,
		eventBroker: broker,
	}

	ss.authService = &AuthService{
		services: ss,
		userModel: models.GetUserModel(),
	}

	ss.userService = &UserService{
		services: ss,
		userModel: models.GetUserModel(),
	}

	return ss, nil
}

func (ss *serviceSupplier) GetEventBroker() event.Broker {
	return ss.eventBroker
}

func (ss *serviceSupplier) GetAuthService() *AuthService {
	return ss.authService
}

func (ss *serviceSupplier) GetUserService() *UserService {
	return ss.userService
}

func (ss *serviceSupplier) release() error {
	return nil
}
