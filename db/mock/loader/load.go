// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package loader

import (
	"github.com/chinmobi/gin-mvc/db/impl"
	db "github.com/chinmobi/gin-mvc/db/mock"
	mock "github.com/chinmobi/gin-mvc/model/mock"
)

func Load(supplier *impl.ModelSupplier) error {
	supplier.SetUserModel(mock.NewUserModel(db.NewUsersDB()))

	return nil
}
