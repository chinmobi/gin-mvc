// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package token

import (
	"github.com/chinmobi/gin-mvc/security"
)

type AbstractAuthenticationToken struct {
	authorities  []security.GrantedAuthority
	details      interface{}
	isAuthed     bool
}

type AbstractAuthToken = AbstractAuthenticationToken

func (aa *AbstractAuthToken) GetAuthorities() []security.GrantedAuthority {
	return aa.authorities
}

//func (aa *AbstractAuthToken) GetCredentials() interface{} {
//}

func (aa *AbstractAuthToken) GetDetails() interface{} {
	return aa.details
}

//func (aa *AbstractAuthToken) GetPrincipal() interface{} {
//}

func (aa *AbstractAuthToken) SetAuthenticated(isAuthenticated bool) {
	aa.isAuthed = isAuthenticated
}

func (aa *AbstractAuthToken) IsAuthenticated() bool {
	return aa.isAuthed
}

// Setter methods

func (aa *AbstractAuthToken) AddAuthority(authority ...security.GrantedAuthority) {
	aa.authorities = append(aa.authorities, authority...)
}

func (aa *AbstractAuthToken) SetDetails(details interface{}) {
	aa.details = details
}
