// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

type SecurityContext interface {
	GetAuthentication() Authentication
	SetAuthentication(auth Authentication) Authentication
	CleanAuthentication() Authentication
}

type Context struct {
	auth Authentication
}

func NewContext() *Context {
	ctx := &Context{
	}
	return ctx;
}

// SecurityContext methods

func (c *Context) GetAuthentication() Authentication {
	return c.auth
}

func (c *Context) SetAuthentication(auth Authentication) Authentication {
	old := c.auth
	c.auth = auth
	return old
}

func (c *Context) CleanAuthentication() Authentication {
	old := c.auth
	c.auth = nil
	return old
}
