// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/auth/token"
	"github.com/chinmobi/gin-mvc/security/web/auth"

	"github.com/gin-gonic/gin"
)

// AnonymousHelper
type AnonymousHelper struct{}

func (a AnonymousHelper) AttemptAuthentication(c *gin.Context, h auth.AuthSuccessHandler) (auth.Authentication, error) {
	token := token.NewAnonymousAuthToken("ANONYMOUS_USER")
	token.AddAuthority(security.SGAuthority("ROLE_ANONYMOUS"))

	token.SetAuthenticated(true)

	// DON'T invoke the AuthSuccessHandler as the AnonymousAuthToken is not a real Authentication.

	return token, nil
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
