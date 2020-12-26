// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/chinmobi/gin-mvc/config"
)

type App struct {
	config  *config.Config
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
	return nil
}

func (app *App) Shutdown() {
	// Shutting down / releasing application components.
}

func (app *App) Config() *config.Config {
	return app.config
}
