// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/security/access"
)

type InterceptorBuilder struct {
	permissions   *access.PermissionsGroup
	deniedHandler AccessDeniedHandler
	evaluator     access.PrivilegeEvaluator
}

func NewBuilder(onAccessDenied OnAccessDeniedFunc) *InterceptorBuilder {
	builder := &InterceptorBuilder{
		deniedHandler: WrapAccessDeniedFunc(onAccessDenied),
		evaluator: access.SimplePrivilegeEvaluator{},
	}

	return builder
}

func (b *InterceptorBuilder) SetPermissions(permissions *access.PermissionsGroup) {
	b.permissions = permissions
}

func (b *InterceptorBuilder) SetDeniedHandler(handler AccessDeniedHandler) {
	b.deniedHandler = handler
}

func (b *InterceptorBuilder) SetEvaluator(evaluator access.PrivilegeEvaluator) {
	b.evaluator = evaluator
}

func (b *InterceptorBuilder) Build() *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    b.permissions,
		deniedHandler:  b.deniedHandler,
		evaluator:      b.evaluator,
	}
	return si
}

func (b *InterceptorBuilder) BuildFor(entry *PermissionsEntry) *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    &entry.permissions,
		deniedHandler:  b.deniedHandler,
		evaluator:      b.evaluator,
	}

	entry.interceptor = si

	return si
}
