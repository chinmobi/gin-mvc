// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

type ControllerSet struct {
}

func NewControllerSet() *ControllerSet {
	set := &ControllerSet{
	}
	return set
}

func (set *ControllerSet) setUp() error {
	return nil
}

func (set *ControllerSet) tearDown() error {
	return nil
}
