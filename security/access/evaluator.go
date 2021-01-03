// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/security"
)

type Authentication = security.Authentication

type PermissionEvaluator interface {
	HasPermission(authentication Authentication, permission RolePermission) bool
}

func SimplePermissionEval(authentication Authentication, permission RolePermission) bool {
	authorities := authentication.GetAuthorities()
	size := len(authorities)
	if size == 0 {
		return false
	}

	permStr := permission.GetPermission()
	for i := 0; i < size; i++ {
		if authorities[i].GetAuthority() == permStr {
			return true
		}
	}

	return false
}

type PrivilegeEvaluator interface {
	IsAllowed(authentication Authentication, permissions *PermissionsGroup) bool
}

func SimplePrivilegeEval(authentication Authentication, permissions *PermissionsGroup) bool {
	authorities := authentication.GetAuthorities()
	authSize := len(authorities)

	rolePermissions := permissions.RolePermissions()
	permSize := len(rolePermissions)

	if authSize == 0 {
		if permSize == 0 {
			return true
		} else {
			return false
		}
	}

	if permSize == 0 {
		return true
	}

	for i := 0; i < permSize; i++ {
		permStr := rolePermissions[i].GetPermission()

		for j := 0; j < authSize; j++ {
			if authorities[j].GetAuthority() == permStr {
				return true
			}
		}
	}

	return false
}

type SimplePrivilegeEvaluator struct{}

func (s SimplePrivilegeEvaluator) IsAllowed(auth Authentication, perms *PermissionsGroup) bool {
	return SimplePrivilegeEval(auth, perms)
}
