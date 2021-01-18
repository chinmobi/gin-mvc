// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/chinmobi/modlib/evt/event"
	"github.com/chinmobi/modlib/grpool/gr"
)

type DefaultMulticaster struct{}

func (m DefaultMulticaster) MulticastEvent(e *event.Event) {
	e.Run()
}

type ExecutorMulticaster struct{
	executor gr.ExecutorService
}

func (m ExecutorMulticaster) MulticastEvent(e *event.Event) {
	err := m.executor.Execute(e)
	if err != nil {
		// TODO
	}
}
