// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/security/access"
)

func buildDefaultInterceptor(entry *PermissionsEntry) *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    &entry.permissions,
		deniedHandler:  NullAccessDeniedHandler{},
		evaluator:      access.SimplePrivilegeEvaluator{},
	}

	entry.interceptor = si

	return si
}
