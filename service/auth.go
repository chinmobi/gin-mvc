// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/gin-mvc/model/x/validator"
	"github.com/chinmobi/gin-mvc/security/auth/errors"
	auth "github.com/chinmobi/gin-mvc/security/auth/principal"
)

type AuthService struct {
	services  *serviceSupplier
	userModel  model.UserModel
}

func (svc *AuthService) CreateAuthProvider() *AuthServiceAuthProvider {
	return newAuthProvider(svc)
}

// UserPrincipalService method
func (svc *AuthService) LoadUserByUsername(username string) (auth.UserPrincipal, error) {
	if uid, ok := validator.ParseUserID(username); ok {
		return svc.LoadUserByUserID(uid)
	}

	if validator.IsNicknameValid(username) {
		return svc.LoadUserByNickname(username)
	}

	return svc.LoadUserByEmail(username)
}

func (svc *AuthService) LoadUserByUserID(uid uint32) (*model.UserDetails, error) {
	entity, err := svc.getUserModel().FindUserByID(uid)
	if err != nil {
		return nil, err
	}
	return svc.loadUserUserDetails(entity)
}

func (svc *AuthService) LoadUserByNickname(nickname string) (*model.UserDetails, error) {
	entity, err := svc.getUserModel().FindUserByNickname(nickname)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.NewErrUsernameNotFound()
	}
	return svc.loadUserUserDetails(entity)
}

func (svc *AuthService) LoadUserByEmail(email string) (*model.UserDetails, error) {
	entity, err := svc.getUserModel().FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if entity == nil {
		return nil, errors.NewErrUsernameNotFound()
	}
	return svc.loadUserUserDetails(entity)
}

func (svc *AuthService) loadUserUserDetails(entity *model.UserEntity) (*model.UserDetails, error) {
	user := model.NewUserDetails(entity)
	if err := svc.loadUserAuthorities(user); err != nil {
		return user, err
	}
	return user, nil
}

func (svc *AuthService) loadUserAuthorities(user *model.UserDetails) error {
	user.AddAuthority(model.SGAuthority("ROLE_USER"))
	return nil
}

func (svc *AuthService) getUserModel() model.UserModel {
	return svc.userModel
}
