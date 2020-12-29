// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package impl

import (
	"github.com/chinmobi/gin-mvc/errors"
	"github.com/chinmobi/gin-mvc/model"
)

// The implementation of the model supplier.
type ModelSupplier struct {
	dbClosers  []Closer
	userModel  model.UserModel
}

func NewModelSupplier() *ModelSupplier {
	ms := &ModelSupplier{
	}
	return ms
}

// Manage DB closers

func (ms *ModelSupplier) AddCloser(c ...Closer) {
	ms.dbClosers = append(ms.dbClosers, c...)
}

func (ms *ModelSupplier) Close() error {
	errs := errors.NewErrWrapErrors()

	for _, closer := range ms.dbClosers {
		if err := closer.Close(); err != nil {
			errs.Wrap(err)
		}
	}

	return errs.AsError()
}

// Get / set models

func (ms *ModelSupplier) SetUserModel(u model.UserModel) {
	ms.userModel = u
}

func (ms *ModelSupplier) GetUserModel() model.UserModel {
	return ms.userModel
}
