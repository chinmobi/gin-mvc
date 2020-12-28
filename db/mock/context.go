// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mock

type Context struct {
	db *UsersDB
}

func NewContext() (*Context, error) {
	ctx := &Context{
		db: newUsersDB(),
	}
	return ctx, nil
}

func (ctx *Context) UsersDB() *UsersDB {
	return nil
}

func (ctx *Context) Close() error {
	// Nothing to do.
	return nil
}
