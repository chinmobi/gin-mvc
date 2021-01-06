// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpool

import (
	"github.com/chinmobi/gin-mvc/config"
	"github.com/chinmobi/gin-mvc/grpool/ants"
	"github.com/chinmobi/gin-mvc/grpool/gr"
)

type antsExecutor struct {
	antsPool *ants.Pool
}

func newAntsExecutor(config *config.Grpool) (*antsExecutor, error) {
	var err error

	exec := &antsExecutor{
	}

	options := configOptions(config)

	exec.antsPool, err = ants.NewPool(config.Size, ants.WithOptions(*options))

	return exec, err
}

func (exec *antsExecutor) Execute(task gr.Runnable) error {
	return exec.antsPool.Submit(task)
}

func (exec *antsExecutor) Shutdown() {
	exec.antsPool.Release()
}

func configOptions(config *config.Grpool) *ants.Options {
	opts := &ants.Options{
		PreAlloc: true,
		Nonblocking: true,
		// TODO
		//Logger: ,
	}

	return opts
}
