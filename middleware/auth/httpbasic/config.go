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

	"github.com/chinmobi/ginmod/security/auth/token"
	"github.com/chinmobi/ginmod/security/web/auth"
	"github.com/chinmobi/gin-mvc/app"

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

func (h HttpBasicHelper) AttemptAuthentication(c *gin.Context) (auth.Authentication, error) {
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

func (h HttpBasicHelper) TearDown() error {
	// Nothing to do.
	return nil
}

func SetHttpBasicRealmHeaderFunc(realm string) auth.OnAuthFailureFunc {
	realm = BASIC_REALM_PREFIX + strconv.Quote(realm)

	return func(c *gin.Context, err *auth.ErrAuthentication) (bool, error) {
		c.Header(AUTHORIZATION_RESP_HEADER, realm)
		return false, nil
	}
}

func Configure(authGroup *auth.ProcessorGroup, app *app.App) error {
	processor := authGroup.CreateProcessor(HttpBasicHelper{})

	// TODO: realm := app.Config.Auth.HttpBasic.Realm
	realm := "Authorization Required"

	authGroup.Configurer().AddAuthFailureFunc(SetHttpBasicRealmHeaderFunc(realm))

	authGroup.AddProcessor(processor)

	return nil
}
