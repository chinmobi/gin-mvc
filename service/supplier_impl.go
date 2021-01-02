// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
)

type serviceSupplier struct {
	models        model.Supplier
	userService  *UserService
}

func createSupplier(models model.Supplier) (*serviceSupplier, error) {
	ss := &serviceSupplier{
		models: models,
	}

	ss.userService = &UserService{
		services: ss,
		userModel: models.GetUserModel(),
	}

	return ss, nil
}

func (ss *serviceSupplier) release() error {
	return nil
}

func (ss *serviceSupplier) GetUserService() *UserService {
	return ss.userService
}
