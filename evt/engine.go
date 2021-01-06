// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package evt

import (
	"github.com/chinmobi/gin-mvc/evt/event"
	"github.com/chinmobi/gin-mvc/evt/internal"
)

// --- publisher ---

type publisher struct {
	topic  string
	engine *engine
}

func newPublisher(topic string, engine *engine) *publisher {
	p := &publisher{
		topic: topic,
		engine: engine,
	}
	return p
}

func (p *publisher) PublishEvent(routingPath, source string, payload event.Payload) {
	p.engineImpl().PublishEvent(p.topic, routingPath, source, payload)
}

func (p *publisher) engineImpl() *internal.Engine {
	return p.engine.engineImpl()
}

// --- engine ---

type engine struct {
	impl *internal.Engine
}

func newEngine(multicaster internal.EventMulticaster) *engine {
	e := &engine{
		impl: internal.NewEngine(multicaster),
	}
	return e
}

// --- Broker methods ---

func (e *engine) Produce(topic string) event.Publisher {
	return newPublisher(topic, e)
}

func (e *engine) Subscribe(topic, bindingPath string, listener event.Listener) {
	e.engineImpl().BindEventListener(topic, bindingPath, listener)
}

func (e *engine) engineImpl() *internal.Engine {
	return e.impl
}
