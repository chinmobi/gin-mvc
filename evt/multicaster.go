// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package evt

import (
	"github.com/chinmobi/gin-mvc/evt/internal"
	"github.com/chinmobi/gin-mvc/grpool/gr"
)

type multicaster struct {
	executor gr.Executor
}

func newMulticaster(executor gr.Executor) *multicaster {
	m := &multicaster{
		executor: executor,
	}
	return m
}

func (m *multicaster) MulticastEvent(event *internal.Event) {
	err := m.executor.Execute(event)
	if err != nil {
		// TODO
	}
}
