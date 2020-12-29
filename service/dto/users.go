// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package dto

import (
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/gin-mvc/model/x/passwd"
)

// NOTE: The user dto is just for demo, could to be removed for real project.

type UserDTO struct {
	Nickname  string  `json:"nickname,omitempty"`
	Password  string  `json:"password,omitempty"`
	Email     string  `json:"email,omitempty"`
}

func (dto *UserDTO) ToUserEntity() (*model.UserEntity, error) {
	u := &model.UserEntity{
		Nickname: dto.Nickname,
		Email: dto.Email,
	}

	if dto.Password != "" {
		hashedPassword, err := passwd.HashPassword(dto.Password)
		if err != nil {
			return nil, err
		}
		u.PasswordHash = string(hashedPassword)
	}

	return u, nil
}
