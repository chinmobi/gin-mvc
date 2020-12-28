// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
)

// NOTE: The user service is just for demo, could to be removed for real project.

type UserService struct {
	services  *serviceSupplier
	userModel  model.UserModel
}
