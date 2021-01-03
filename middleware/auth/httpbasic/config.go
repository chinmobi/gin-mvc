// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package httpbasic

import (
	"bytes"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/security/auth/token"
	"github.com/chinmobi/gin-mvc/security/web/auth"

	"github.com/gin-gonic/gin"
)

const (
	AUTHORIZATION_REQ_HEADER  = "Authorization"
	AUTHORIZATION_RESP_HEADER = "WWW-Authenticate"
	BASIC_PREFIX              = "Basic "
	BASIC_REALM_PREFIX        = "Basic realm="
)

// HttpBasicHelper
type HttpBasicHelper struct{}

func (hb HttpBasicHelper) AttemptAuthentication(c *gin.Context, h auth.AuthSuccessHandler) (auth.Authentication, error) {
	authHeader := c.GetHeader(AUTHORIZATION_REQ_HEADER)
	if authHeader == "" {
		return nil, nil
	}

	if !strings.HasPrefix(authHeader, BASIC_PREFIX) {
		return nil, nil
	}

	authValue := authHeader[len(BASIC_PREFIX):]

	namePass, err := base64.StdEncoding.DecodeString(authValue)
	if err != nil {
		return nil, err
	}

	index := bytes.LastIndexByte(namePass, byte(':'))
	if index < 0 {
		return nil, errors.New("Malformed Authorization header Basic value")
	}

	username := string(namePass[:index])
	password := string(namePass[index+1:])

	return token.NewUsernamePasswordAuthToken(username, password), nil
}

func (hb HttpBasicHelper) TearDown() error {
	// Nothing to do.
	return nil
}

func SetHttpBasicRealmHeaderFunc(realm string) auth.OnAuthFailureFunc {
	realm = BASIC_REALM_PREFIX + strconv.Quote(realm)

	return func(c *gin.Context, err *auth.ErrAuthentication) error {
		c.Header(AUTHORIZATION_RESP_HEADER, realm)
		return nil
	}
}

func Configure(authGroup *auth.ProcessorGroup, app *app.App) error {
	processor := authGroup.CreateProcessor(HttpBasicHelper{})

	// No AuthProvider for httpbasic processor.
	//authGroup.Configurer().AddProvider(...)

	// No OnAuthSuccessFunc to do for httpbasic processor.
	//authGroup.Configurer().AddSuccessFunc(...)

	// TODO: realm := app.Config.Auth.HttpBasic.Realm
	realm := "Authorization Required"

	authGroup.Configurer().AddFailureFunc(SetHttpBasicRealmHeaderFunc(realm))

	authGroup.AddProcessor(processor)

	return nil
}
