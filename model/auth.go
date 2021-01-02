// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"strconv"

	"github.com/chinmobi/gin-mvc/model/x/passwd"
	auth "github.com/chinmobi/gin-mvc/security/auth/principal"
)

type SGAuthority = auth.SGAuthority

type UserDetails struct {
	authorities  []auth.GrantedAuthority
	entity       *UserEntity
	username     string
	details      interface{}
}

func NewUserDetails(user *UserEntity) *UserDetails {
	u := &UserDetails{
		authorities: []auth.GrantedAuthority{},
		entity: user,
		username: strconv.FormatUint(uint64(user.ID), 10),
	}
	return u
}

func (u *UserDetails) SetDetails(details interface{}) {
	u.details = details
}

func (u *UserDetails) AddAuthority(authority ...auth.GrantedAuthority) {
	u.authorities = append(u.authorities, authority...)
}

// UserPrincipal methods

func (u *UserDetails) GetAuthorities() []auth.GrantedAuthority {
	return u.authorities
}

func (u *UserDetails) GetUsername() string {
	return u.username
}

func (u *UserDetails) GetPassword() string {
	return u.entity.PasswordHash
}

func (u *UserDetails) IsPasswordMatched(password string) bool {
	passwordHash := u.GetPassword()
	if password == "" || passwordHash == "" {
		return false
	}
	if err := passwd.VerifyPassword(passwordHash, password); err != nil {
		return false
	}
	return true
}

func (u *UserDetails) GetDetails() interface{} {
	return u.details
}

func (u *UserDetails) IsAccountNonExpired() bool {
	return true
}

func (u *UserDetails) IsAccountNonLocked() bool {
	return true
}

func (u *UserDetails) IsCredentialsNonExpired() bool {
	return true
}

func (u *UserDetails) IsEnabled() bool {
	return true
}

// UserDetails's Authentication token

type UserDetailsToken = auth.UserPrincipalToken

func NewUserDetailsToken(user *UserDetails) *UserDetailsToken {
	return auth.NewUserPrincipalToken(user)
}
