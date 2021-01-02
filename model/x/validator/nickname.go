// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validator

import (
	"regexp"
	"strings"
)

const (
	nicknameExp = "^[a-zA-Z][a-zA-Z0-9._-]{3,23}$"
)

var nicknameRegexp = regexp.MustCompile(nicknameExp)

func IsNicknameValid(n string) bool {
	if strings.IndexByte(n, '@') >= 0 {
		return false
	}
	return nicknameRegexp.MatchString(n)
}
