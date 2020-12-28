// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/db/impl"
	"github.com/chinmobi/gin-mvc/db/loader"
	"github.com/chinmobi/gin-mvc/model"
)

func Load(config *config.Config) (model.Supplier, error) {
	supplier := impl.NewModelSupplier()

	return loader.Load(config, supplier)
}

func Release(supplier model.Supplier) error {
	if ms, ok := supplier.(*impl.ModelSupplier); ok {
		return ms.Close()
	}
	return nil
}
