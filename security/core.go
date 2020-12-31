// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

import (
	"strings"
)

const ROLE_PREFIX = "ROLE_"

type GrantedAuthority interface {
	GetAuthority() string
	IsRole() bool
}

type Authentication interface {
	GetAuthorities()  []GrantedAuthority

	GetCredentials()  interface{}
	GetDetails()      interface{}

	GetPrincipal()    interface{}

	SetAuthenticated(isAuthenticated bool)
	IsAuthenticated() bool
}

// SimpleGrantedAuthority
type SimpleGrantedAuthority string

func (s SimpleGrantedAuthority) GetAuthority() string {
	return string(s)
}

func (s SimpleGrantedAuthority) IsRole() bool {
	return strings.HasPrefix(string(s), ROLE_PREFIX)
}

type SGAuthority = SimpleGrantedAuthority
