// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mock

// NOTE: The mock user model is just for demo, could to be removed for real project.

import (
	"github.com/chinmobi/gin-mvc/db/mock"
	"github.com/chinmobi/gin-mvc/model"
)

type Users struct {
	db *mock.UsersDB
}

func NewUserModel(db *mock.UsersDB) *Users {
	users := &Users{
		db: db,
	}
	return users
}

func (users *Users) CreateUser(u *model.UserEntity) (*model.UserEntity, error) {
	return users.db.CreateUser(u)
}

func (users *Users) FindAllUsers() ([]model.UserEntity, error) {
	return users.db.FindAllUsers()
}

func (users *Users) FindUserByID(uid uint32) (*model.UserEntity, error) {
	return users.db.FindUserByID(uid)
}

func (users *Users) FindUserByNickname(nickname string) (*model.UserEntity, error) {
	return users.db.FindUserByNickname(nickname)
}

func (users *Users) FindUserByEmail(email string) (*model.UserEntity, error) {
	return users.db.FindUserByEmail(email)
}

func (users *Users) UpdateUser(uid uint32, u *model.UserEntity) (*model.UserEntity, error) {
	return users.db.UpdateUser(uid, u)
}

func (users *Users) DeleteUser(uid uint32) (uint64, error) {
	return users.db.DeleteUser(uid)
}
