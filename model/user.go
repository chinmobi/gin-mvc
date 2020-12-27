// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

// NOTE: The user model is just for demo, could to be removed for real project.

type UserEntity struct {
	ID           uint32  `json:"id,omitempty"`
	Nickname     string  `json:"nickname,omitempty"`
	PasswordHash string  `json:"password,omitempty"`
	Email        string  `json:"email,omitempty"`
}

func (u *UserEntity) SameIdentityAs(other *UserEntity) bool {
	return u.ID == other.ID
}

type UserModel interface {
	CreateUser(u *UserEntity) (*UserEntity, error)
	FindAllUsers() ([]UserEntity, error)
	FindUserByID(uid uint32) (*UserEntity, error)
	FindUserByNickname(nickname string) (*UserEntity, error)
	FindUserByEmail(email string) (*UserEntity, error)
	UpdateUser(uid uint32, u *UserEntity) (*UserEntity, error)
	DeleteUser(uid uint32) (uint64, error)
}
