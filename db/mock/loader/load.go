// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package loader

import (
	"github.com/chinmobi/gin-mvc/model"
	mock "github.com/chinmobi/gin-mvc/model/mock"
	db "github.com/chinmobi/gin-mvc/db/mock"
)

func Load(set *model.SupplierSet) error {
	set.SetUserModel(mock.NewUserModel(db.NewUsersDB()))

	return nil
}
