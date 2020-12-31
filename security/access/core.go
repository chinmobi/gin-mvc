// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"strings"

	"github.com/chinmobi/gin-mvc/security"
)

const ROLE_PREFIX = security.ROLE_PREFIX

type NeedPermission interface {
	GetPermission() string
	IsRole() bool
}

// SimpleNeedPermission
type SimpleNeedPermission string

func (s SimpleNeedPermission) GetPermission() string {
	return string(s)
}

func (s SimpleNeedPermission) IsRole() bool {
	return strings.HasPrefix(string(s), ROLE_PREFIX)
}

type SNPermission = SimpleNeedPermission

// Permissions' group
type PermissionsGroup struct {
	name         string
	permissions  []NeedPermission
}

func NewPermissionsGroup(name string) *PermissionsGroup {
	pg := &PermissionsGroup{
		name: name,
	}
	return pg
}

func (pg *PermissionsGroup) AddPermission(permission ...NeedPermission) {
	pg.permissions = append(pg.permissions, permission...)
}

func (pg *PermissionsGroup) NeedPermissions() []NeedPermission {
	return pg.permissions
}

func (pg *PermissionsGroup) Name() string {
	return pg.name
}
