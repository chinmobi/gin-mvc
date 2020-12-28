// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/db"
	"github.com/chinmobi/gin-mvc/errors"
	"github.com/chinmobi/gin-mvc/model"
)

type App struct {
	config        *config.Config
	modelSupplier  model.Supplier
}

func New(config *config.Config) *App {
	app := &App{
		config: config,
	}

	return app
}

func NewWithStart(config *config.Config) (*App, error) {
	app := New(config)

	return app, app.Start()
}

func (app *App) Start() error {
	// Configuring, setting up / starting application components.

	modelSupplier, err := db.Load(app.config)
	if err != nil {
		return err
	}
	app.modelSupplier = modelSupplier

	return nil
}

func (app *App) Shutdown() error {
	// Shutting down / releasing application components.

	var errs *errors.ErrWrapErrors

	if err := db.Release(app.modelSupplier); err != nil {
		if errs == nil {
			errs = errors.NewErrWrapErrors()
		}
		errs.Wrap(err)
	}

	return errs
}

func (app *App) Config() *config.Config {
	return app.config
}

func (app *App) ModelSupplier() model.Supplier {
	return app.modelSupplier
}
