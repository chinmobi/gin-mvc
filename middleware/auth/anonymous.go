// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	sec "github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/auth"

	"github.com/gin-gonic/gin"
)

// AnonymousToken
type AnonymousToken struct{}

func (a AnonymousToken) GetAuthorities() []sec.GrantedAuthority {
	return []sec.GrantedAuthority{ sec.SGAuthority("ROLE_ANONYMOUS") }
}

func (a AnonymousToken) GetCredentials() interface{} {
	return ""
}

func (a AnonymousToken) GetDetails() interface{} {
	return ""
}

func (a AnonymousToken) GetPrincipal() interface{} {
	return "ANONYMOUS_USER"
}

func (a AnonymousToken) SetAuthenticated(isAuthenticated bool) {
	// Nothing to do.
}

func (a AnonymousToken) IsAuthenticated() bool {
	return true
}

// AnonymousHelper
type AnonymousHelper struct{}

func (a AnonymousHelper) AttemptAuthentication(c *gin.Context, h sec.AuthSuccessHandler) (sec.Authentication, error) {
	// DON'T invoke the AuthSuccessHandler as the AnonymousToken is not a real Authentication.
	return AnonymousToken{}, nil
}

func (a AnonymousHelper) TearDown() error {
	// Nothing to do.
	return nil
}

// Configure the anonymous processor

func configureAnonymous(authGroup *auth.ProcessorGroup) {
	processor := authGroup.CreateProcessor(AnonymousHelper{})

	// No AuthProvider for anonymous processor.
	//authGroup.Configurer().AddProvider(...)

	// No OnAuthSuccessFunc to do for anonymous processor.
	//authGroup.Configurer().AddSuccessFunc(...)

	authGroup.AddProcessor(processor)
}
