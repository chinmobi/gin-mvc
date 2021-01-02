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

type UserPrincipalAuthToken struct {
	user UserPrincipal
	isAuthed bool
}

func NewUserPrincipalAuthToken(u UserPrincipal) *UserPrincipalAuthToken {
	token := &UserPrincipalAuthToken{
		user: u,
		isAuthed: false,
	}
	return token
}

func (u *UserPrincipalAuthToken) GetAuthorities() []GrantedAuthority {
	return u.user.GetAuthorities()
}

func (u *UserPrincipalAuthToken) GetCredentials() interface{} {
	return u.user.GetPassword()
}

func (u *UserPrincipalAuthToken) GetDetails() interface{} {
	return u.user.GetDetails()
}

func (u *UserPrincipalAuthToken) GetPrincipal() interface{} {
	return u.user
}

func (u *UserPrincipalAuthToken) SetAuthenticated(isAuthenticated bool) {
	u.isAuthed = isAuthenticated
}

func (u *UserPrincipalAuthToken) IsAuthenticated() bool {
	return u.isAuthed
}
