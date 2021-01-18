// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"github.com/chinmobi/modlib/evt/event"
)

type DefaultMulticaster struct{}

func (m DefaultMulticaster) MulticastEvent(e *event.Event) {
	e.Run()
}
