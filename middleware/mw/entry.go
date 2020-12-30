// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mw

const (
	MW_COMMON = "MW_COMMON"
	MW_AUTH   = "MW_AUTH"
)

type Entry struct {
	category  string
	adapters  []Adapter
	next      *Entry
}

func NewEntry(category string) *Entry {
	e := &Entry{
		category: category,
	}
	return e
}

func (e *Entry) NewNext(category string) *Entry {
	n := &Entry{
		category: category,
	}
	e.next = n
	return n
}

func (e *Entry) Category() string {
	return e.category
}

func (e *Entry) Next() *Entry {
	return e.next
}

func (e *Entry) putAdapter(adapter ...Adapter) {
	e.adapters = append(e.adapters, adapter...)
}

func (e *Entry) Adapters() []Adapter {
	return e.adapters
}
