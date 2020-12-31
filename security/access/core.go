// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"strings"

	"github.com/chinmobi/gin-mvc/security"
)

const ROLE_PREFIX = security.ROLE_PREFIX

type RolePermission interface {
	GetPermission() string
	IsRole() bool
}

// SimpleRolePermission

type SimpleRolePermission string

func (s SimpleRolePermission) GetPermission() string {
	return string(s)
}

func (s SimpleRolePermission) IsRole() bool {
	return strings.HasPrefix(string(s), ROLE_PREFIX)
}

type SRPermission = SimpleRolePermission

// PermissionsGroup

type PermissionsGroup struct {
	name         string
	permissions  []RolePermission
}

func NewPermissionsGroup(name string) *PermissionsGroup {
	pg := &PermissionsGroup{
		name: name,
	}
	return pg
}

func (pg *PermissionsGroup) Clear() {
	if pg.permissions != nil {
		pg.permissions = pg.permissions[0:0]
	}
}

func (pg *PermissionsGroup) AddPermission(permission ...RolePermission) {
	pg.permissions = append(pg.permissions, permission...)
}

func (pg *PermissionsGroup) RolePermissions() []RolePermission {
	return pg.permissions
}

func (pg *PermissionsGroup) Name() string {
	return pg.name
}
