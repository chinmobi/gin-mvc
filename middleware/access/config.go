// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/gin-mvc/app"
	"github.com/chinmobi/gin-mvc/security/access"
)

type SRPermission = access.SimpleRolePermission

func Configure(config *access.PermissionsConfigurer, app *app.App) error {
	// Configure the role permissions

	entry := config.ConfigureEntry("CUD_USERS")

	entry.AddPermission(SRPermission("USERS_C_U_D"))
	entry.AddPermission(SRPermission("ROLE_ADMIN"))

	// Build interceptor for entry

	interceptorBuilder := setUpBuilder()

	interceptorBuilder.BuildFor(entry)

	return nil
}
