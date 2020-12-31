// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/gin-gonic/gin"
)

const CTX_ACCESS_PERMISSIONS = "CTX_ACCESS_PERMISSIONS"

// PermissionsEntry

type PermissionsEntry struct {
	permissions  PermissionsGroup
	handlerFunc  gin.HandlerFunc
	interceptor  *SecurityInterceptor
}

func (pe *PermissionsEntry) Reset() {
	pe.permissions.Clear()
	pe.handlerFunc = nil
	pe.interceptor = nil
}

func (pe *PermissionsEntry) PermissionsGroup() *PermissionsGroup {
	return &pe.permissions
}

func (pe *PermissionsEntry) AddPermission(permission ...RolePermission) {
	pe.permissions.AddPermission(permission...)
}

func (pe *PermissionsEntry) configure(c *gin.Context) {
	c.Set(CTX_ACCESS_PERMISSIONS, &pe.permissions)

	c.Next()
}

func (pe *PermissionsEntry) ConfigureHandlerFunc() gin.HandlerFunc {
	if pe.handlerFunc == nil {
		pe.handlerFunc = pe.createHandlerFunc()
	}
	return pe.handlerFunc
}

func (pe *PermissionsEntry) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		pe.configure(c)
	}
}

func (pe *PermissionsEntry) AccessDecisionAgent() *AccessDecisionAgent {
	if pe.interceptor == nil {
		pe.interceptor = buildDefaultInterceptor(pe)
	}

	return pe.interceptor.DecisionAgent()
}

// PermissionsConfigurer

type PermissionsConfigurer struct {
	entries []*PermissionsEntry
}

func NewPermissionsConfigurer() *PermissionsConfigurer {
	pc := &PermissionsConfigurer{
	}
	return pc
}

func (pc *PermissionsConfigurer) Reset() {
	if pc.entries != nil {
		for i, cnt := 0, len(pc.entries); i < cnt; i++ {
			pc.entries[i].Reset()
		}

		pc.entries = pc.entries[0:0]
	}
}

func (pc *PermissionsConfigurer) ConfigureEntry(group string) *PermissionsEntry {
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
