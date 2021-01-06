// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

import (
	"github.com/chinmobi/gin-mvc/evt/event"
)

type EventPayload = event.Payload

type EventListener = event.Listener

// --- EventHandler ---

type EventHandler interface {
	HandleEvent(event *Event)
}

// --- Event ---

type Event struct {
	Topic, RoutingPath, Source string
	Payload EventPayload
	Handler EventHandler
}

func NewEvent() *Event {
	event := &Event{
	}
	return event
}

func (e *Event) Init(topic, routingPath, source string, payload EventPayload) {
	e.Topic = topic
	e.RoutingPath = routingPath
	e.Source = source
	e.Payload = payload
}

func (e *Event) Reset() *Event {
	e.Topic = ""
	e.RoutingPath = ""
	e.Source = ""
	e.Payload = nil

	e.Handler = nil

	return e
}

func (e *Event) Run() {
	if e.Handler != nil {
		e.Handler.HandleEvent(e)
	}
}
