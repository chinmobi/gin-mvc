// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package loader

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/db/impl"
	mock "github.com/chinmobi/gin-mvc/db/mock/loader"
	"github.com/chinmobi/gin-mvc/model"
)

func Load(config *config.Config, supplier *impl.ModelSupplier) (model.Supplier, error) {
	if err := mock.Load(supplier); err != nil {
		return supplier, err
	}

	return supplier, nil
}
