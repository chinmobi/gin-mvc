// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
)

type Supplier interface {
	GetUserService() *UserService
}

func SetUp(models model.Supplier) (Supplier, error) {
	return createSupplier(models)
}

func TearDown(services Supplier) error {
	if ss, ok := services.(*serviceSupplier); ok {
		return ss.release()
	}
	return nil
}
