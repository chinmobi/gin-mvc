// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/chinmobi/gin-mvc/model"
)

// The implementation of the model supplier.
type ModelSupplier struct {
	userModel model.UserModel
}

func NewModelSupplier() *ModelSupplier {
	ms := &ModelSupplier{
	}
	return ms
}

func (ms *ModelSupplier) SetUserModel(u model.UserModel) *ModelSupplier {
	ms.userModel = u
	return ms
}

func (ms *ModelSupplier) GetUserModel() model.UserModel {
	return ms.userModel
}
