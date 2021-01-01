// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/errors"
	acconfig "github.com/chinmobi/gin-mvc/middleware/access"
	"github.com/chinmobi/gin-mvc/middleware/config"
	"github.com/chinmobi/gin-mvc/middleware/mw"
	"github.com/chinmobi/gin-mvc/security"
	"github.com/chinmobi/gin-mvc/security/access"
	"github.com/chinmobi/gin-mvc/security/auth"

	"github.com/gin-gonic/gin"
)

type MiddlewareSet struct {
	authConfig   *auth.ProcessorConfigurer
	permsConfig  *access.PermissionsConfigurer
	entries      *mw.Entry
}

func NewMiddlewareSet() *MiddlewareSet {
	set := &MiddlewareSet{
		authConfig: auth.NewProcessorConfigurer(),
		permsConfig: access.NewPermissionsConfigurer(),
	}
	return set
}

func (set *MiddlewareSet) setUp(app *app.App) error {
	set.entries = mw.NewEntry(mw.MW_COMMON)
	set.entries.NewNext(mw.MW_AUTH)

	configurer := mw.NewConfigurer(set.entries)

	if err := config.Configure(configurer, set.authConfig, app); err != nil {
		set.tearDown()
		return err
	}

	if err := acconfig.Configure(set.permsConfig, app); err != nil {
		set.tearDown()
		return err
	}

	return nil
}

func (set *MiddlewareSet) tearDown() error {
	errs := errors.NewErrWrapErrors()

	set.permsConfig.Reset()

	entry := set.entries
	for entry != nil {
		adapters := entry.Adapters()
		for i := len(adapters)-1; i >= 0; i-- {
			if err := adapters[i].TearDown(); err != nil {
				errs.Wrap(err)
			}
		}

		entry = entry.Next()
	}

	set.authConfig.Reset()

	return errs.AsError()
}

func (set *MiddlewareSet) SecurityAuthHandler() security.AuthHandler {
	return set.authConfig.AuthHandlerSet()
}

func (set *MiddlewareSet) CommonHandlersChain() gin.HandlersChain {
	return set.HandlersChain(mw.MW_COMMON)
}

func (set *MiddlewareSet) AuthHandlersChain() gin.HandlersChain {
	return set.HandlersChain(mw.MW_AUTH)
}

func (set *MiddlewareSet) PermissionsConfigurer() *access.PermissionsConfigurer {
	return set.permsConfig
}

func (set *MiddlewareSet) HandlersChain(category string) gin.HandlersChain {
	handlers := gin.HandlersChain{}

	entry := set.findEntry(category)
	if entry == nil {
		return handlers
	}

	adapters := entry.Adapters()
	for i, cnt := 0, len(adapters); i < cnt; i++ {
		handlers = append(handlers, adapters[i].HandlerFunc())
	}

	return handlers
}

func (set *MiddlewareSet) findEntry(category string) *mw.Entry {
	entry := set.entries
	for entry != nil {
		if entry.Category() == category {
			return entry
		}

		entry = entry.Next()
	}
	return entry
}
