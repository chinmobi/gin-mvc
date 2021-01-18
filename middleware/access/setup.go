// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/ginmod/security/web/access"
)

func setUpBuilder() *access.InterceptorBuilder {
	builder := access.NewBuilder(RespAccessDenied)

	return builder
}
