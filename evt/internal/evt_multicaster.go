// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package internal

type EventMulticaster interface {
	MulticastEvent(event *Event)
}
