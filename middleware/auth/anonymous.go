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

func (a AnonymousHelper) AttemptAuthentication(c *gin.Context) (auth.Authentication, error) {
	token := token.NewAnonymousAuthToken("ANONYMOUS_USER")
	token.AddAuthority(security.SGAuthority("ROLE_ANONYMOUS"))

	return token, nil
}

func (a AnonymousHelper) TearDown() error {
	// Nothing to do.
	return nil
}

// Configure the anonymous processor

func configureAnonymous(authGroup *auth.ProcessorGroup) {
	processor := authGroup.CreateProcessor(AnonymousHelper{})

	// Need AuthProvider for anonymous processor.
	//authGroup.Configurer().AddProvider(...)

	authGroup.AddProcessor(processor)
}
