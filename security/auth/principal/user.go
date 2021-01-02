// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package principal

import (
	"github.com/chinmobi/gin-mvc/security"
)

type GrantedAuthority = security.GrantedAuthority

type SimpleGrantedAuthority = security.SimpleGrantedAuthority
type SGAuthority = security.SimpleGrantedAuthority

// UserPrincipal
type UserPrincipal interface {
	GetAuthorities() []GrantedAuthority

	GetUsername() string

	GetPassword() string

	IsPasswordMatched(password string) bool

	GetDetails() interface{}

	IsAccountNonExpired() bool
	IsAccountNonLocked() bool
	IsCredentialsNonExpired() bool
	IsEnabled() bool
}

// UserPrincipalService
type UserPrincipalService interface {
	LoadUserByUsername(username string) (UserPrincipal, error)
}

// UserPrincipal's Authentication token

type UserPrincipalToken struct {
	user UserPrincipal
	isAuthed bool
}

func NewUserPrincipalToken(u UserPrincipal) *UserPrincipalToken {
	token := &UserPrincipalToken{
		user: u,
		isAuthed: false,
	}
	return token
}

func (u *UserPrincipalToken) GetAuthorities() []GrantedAuthority {
	return u.user.GetAuthorities()
}

func (u *UserPrincipalToken) GetCredentials() interface{} {
	return u.user.GetPassword()
}

func (u *UserPrincipalToken) GetDetails() interface{} {
	return u.user.GetDetails()
}

func (u *UserPrincipalToken) GetPrincipal() interface{} {
	return u.user
}

func (u *UserPrincipalToken) SetAuthenticated(isAuthenticated bool) {
	u.isAuthed = isAuthenticated
}

func (u *UserPrincipalToken) IsAuthenticated() bool {
	return u.isAuthed
}
