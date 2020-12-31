// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/gin-gonic/gin"
)

const CTX_ACCESS_PERMISSIONS = "CTX_ACCESS_PERMISSIONS"

type PermissionsEntry struct {
	permissions  PermissionsGroup
	handlerFunc  gin.HandlerFunc
	interceptor  *SecurityInterceptor
}

func (pe *PermissionsEntry) PermissionsGroup() *PermissionsGroup {
	return &pe.permissions
}

func (pe *PermissionsEntry) handle(c *gin.Context) {
	c.Set(CTX_ACCESS_PERMISSIONS, &pe.permissions)

	c.Next()
}

func (pe *PermissionsEntry) GetHandlerFunc() gin.HandlerFunc {
	if pe.handlerFunc == nil {
		pe.handlerFunc = pe.doGetHandlerFunc()
	}
	return pe.handlerFunc
}

func (pe *PermissionsEntry) doGetHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		pe.handle(c)
	}
}

type PermissionsConfigurer struct {
	entries []*PermissionsEntry
}

func (pc *PermissionsConfigurer) Configure(group string) *PermissionsEntry {
	for i, cnt := 0, len(pc.entries); i < cnt; i++ {
		entry := pc.entries[i]
		if entry.permissions.Name() == group {
			return entry
		}
	}

	entry := &PermissionsEntry{
		permissions: PermissionsGroup{
			name: group,
		},
	}

	pc.entries = append(pc.entries, entry)

	return entry
}
