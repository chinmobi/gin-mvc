// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/gin-mvc/service/dto"
)

// NOTE: The user service is just for demo, could to be removed for real project.

type UserService struct {
	services  *serviceSupplier
	userModel  model.UserModel
}

func (svc *UserService) CreateUser(u *dto.UserDTO) (*model.UserEntity, error) {
	entity, err := u.ToUserEntity()
	if err != nil {
		return nil, err
	}
	return svc.getUserModel().CreateUser(entity)
}

func (svc *UserService) FindAllUsers() ([]model.UserEntity, error) {
	return svc.getUserModel().FindAllUsers()
}

func (svc *UserService) FindUserByID(uid uint32) (*model.UserEntity, error) {
	return svc.getUserModel().FindUserByID(uid)
}

func (svc *UserService) UpdateUser(uid uint32, u *dto.UserDTO) (*model.UserEntity, error) {
	entity, err := u.ToUserEntity()
	if err != nil {
		return nil, err
	}
	return svc.getUserModel().UpdateUser(uid, entity)
}

func (svc *UserService) DeleteUser(uid uint32) (uint64, error) {
	return svc.getUserModel().DeleteUser(uid)
}

func (svc *UserService) getUserModel() model.UserModel {
	return svc.userModel
}
