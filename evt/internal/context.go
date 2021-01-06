// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// at https://github.com/gin-gonic/gin/blob/master/LICENSE

package internal

import (
	"math"
)

const abortIndex int8 = math.MaxInt8 / 2

type Context struct {
	Params   Params

	handlers HandlersChain
	index    int8
	fullPath string

	event    *Event

	params   *Params
}

func newContext(maxParams uint16) *Context {
	v := make(Params, 0, maxParams)
	c := &Context{
		params: &v,
	}
	return c
}

func (c *Context) reset() {
	c.Params = c.Params[0:0]

	c.handlers = nil
	c.index = -1
	c.fullPath = ""

	c.event = nil

	*c.params = (*c.params)[0:0]
}

func (c *Context) resetWithEvent(e *Event) {
	c.reset()
	c.event = e
}

func (c *Context) GetEvent() *Event {
	return c.event
}

func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

// Next should be used only inside middleware.
// It executes the pending handlers in the chain inside the calling handler.
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called.
// Note that this will not stop the current handler.
func (c *Context) Abort() {
	c.index = abortIndex
}

// --- Event Envelope methods ---

func (c *Context) GetEventPayload() EventPayload {
	return c.event.Payload
}

func (c *Context) Topic() string {
	return c.event.Topic
}

func (c *Context) RoutingPath() string {
	return c.event.RoutingPath
}

func (c *Context) Source() string {
	return c.event.Source
}

func (c *Context) GetParam(name string) string {
	return c.Param(name)
}

func (c *Context) getEngine() *Engine {
	if c.event.Handler == nil {
		return nil
	}

	engine, ok := c.event.Handler.(*Engine)
	if !ok {
		return nil
	}

	return engine
}

func (c *Context) Reply(ack EventPayload) {
	engine := c.getEngine()
	if engine != nil {
		engine.ReplyEvent(c.event, ack)
	}
}
