// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validator

import (
	"strconv"
	"unicode"
)

func ParseUserID(s string) (uint32, bool) {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return 0, false
		}
	}
	id, _ := strconv.ParseUint(s, 10, 64)
	return uint32(id), true
}
