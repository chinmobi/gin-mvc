// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

type GrantedAuthority interface {
	GetAuthority() string
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

type SGAuthority = SimpleGrantedAuthority
