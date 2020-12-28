// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

// The model supplier provides all of the models that will be used for services.
type Supplier interface {
	GetUserModel() UserModel
}
