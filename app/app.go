// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/db"
	"github.com/chinmobi/gin-mvc/errors"
	"github.com/chinmobi/gin-mvc/evt"
	"github.com/chinmobi/gin-mvc/evt/event"
	"github.com/chinmobi/gin-mvc/grpool"
	"github.com/chinmobi/gin-mvc/grpool/gr"
	"github.com/chinmobi/gin-mvc/log"
	"github.com/chinmobi/gin-mvc/model"
	"github.com/chinmobi/gin-mvc/service"
)

type App struct {
	config      *config.Config
	executor    gr.ExecutorService
	eventBroker event.Broker
	models      model.Supplier
	services    service.Supplier
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

	log.SetUp(&app.config.Logger)
	defer log.L().Sync()

	executor, err := grpool.SetUp(&app.config.Grpool)
	if err != nil {
		app.Shutdown()
		return err
	}
	app.executor = executor

	eventBroker, err := evt.SetUp(executor)
	if err != nil {
		app.Shutdown()
		return err
	}
	app.eventBroker = eventBroker

	// Load the models
	models, err := db.Load(app.config)
	if err != nil {
		app.Shutdown()
		return err
	}
	app.models = models

	// Set up the services
	services, err := service.SetUp(models, eventBroker)
	if err != nil {
		app.Shutdown()
		return err
	}
	app.services = services

	return nil
}

func (app *App) Shutdown() error {
	// Shutting down / releasing application components.

	defer log.L().Sync()

	errs := errors.NewErrWrapErrors()

	if app.executor != nil {
		app.executor.Shutdown()
	}

	// Tear down the services
	if app.services != nil {
		if err := service.TearDown(app.services); err != nil {
			errs.Wrap(err)
		}
		app.services = nil
	}

	// Release the models
	if app.models != nil {
		if err := db.Release(app.models); err != nil {
			errs.Wrap(err)
		}
		app.models = nil
	}

	return errs.AsError()
}

func (app *App) Config() *config.Config {
	return app.config
}

func (app *App) EventBroker() event.Broker {
	return app.eventBroker
}

func (app *App) Executor() gr.Executor {
	return app.executor
}

func (app *App) ServiceSupplier() service.Supplier {
	return app.services
}

func (app *App) ModelSupplier() model.Supplier {
	return app.models
}
