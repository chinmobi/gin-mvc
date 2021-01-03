// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"github.com/chinmobi/gin-mvc/model"
	sec "github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/auth/errors"
	"github.com/chinmobi/gin-mvc/security/auth/token"
)

type AuthServiceAuthProvider struct {
	authService *AuthService
}

func newAuthProvider(svc *AuthService) *AuthServiceAuthProvider {
	provider := &AuthServiceAuthProvider{
		authService: svc,
	}
	return provider
}

// AuthenticationProvider methods
func (p *AuthServiceAuthProvider) Authenticate(auth sec.Authentication) (sec.Authentication, error) {
	if _, ok := auth.(*token.UsernamePasswordAuthToken); ok {
		return p.authUsernamePassword(auth)
	}
	return auth, nil
}

func (p *AuthServiceAuthProvider) Supports(auth sec.Authentication) bool {
	if _, ok := auth.(*token.UsernamePasswordAuthToken); ok {
		return true
	}
	return false
}

func (p *AuthServiceAuthProvider) authUsernamePassword(auth sec.Authentication) (sec.Authentication, error) {
	token := auth.(*token.UsernamePasswordAuthToken)

	principal, err := p.authService.LoadUserByUsername(token.GetUsername())
	if err != nil {
		return token, err
	}

	if !principal.IsPasswordMatched(token.GetPassword()) {
		return token, errors.NewErrBadCredentials()
	}

	if err := checkAccountStatus(principal); err != nil {
		return token, err
	}

	userDetails := principal.(*model.UserDetails)
	userDetailsToken := model.NewUserDetailsAuthToken(userDetails)

	userDetailsToken.SetAuthenticated(true)

	return userDetailsToken, nil
}

func checkAccountStatus(principal model.UserPrincipal) error {
	if !principal.IsAccountNonExpired() {
		return errors.NewErrAccountExpired()
	}
	if !principal.IsAccountNonLocked() {
		return errors.NewErrAccountLocked()
	}
	if !principal.IsCredentialsNonExpired() {
		return errors.NewErrCredentialExpired()
	}
	if !principal.IsEnabled() {
		return errors.NewErrAccountDisabled()
	}
	return nil
}
