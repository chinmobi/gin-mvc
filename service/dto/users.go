// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dto

import (
	"github.com/chinmobi/gin-mvc/model"
)

// NOTE: The user dto is just for demo, could to be removed for real project.

type UserDTO struct {
	Nickname  string  `json:"nickname,omitempty"`
	Password  string  `json:"password,omitempty"`
	Email     string  `json:"email,omitempty"`
}

func (dto *UserDTO) ToUserEntity() *model.UserEntity {
	u := &model.UserEntity{
		Nickname: dto.Nickname,
		PasswordHash: dto.Password,
		Email: dto.Email,
	}
	return u
}
