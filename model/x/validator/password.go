// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package validator

import (
	//"regexp"
	"strings"
	"unicode"
)

const (
	//passwordExp = "^(?=.*[A-Z].*[A-Z])(?=.*[a-z].*[a-z].*[a-z])(?=.*[0-9].*[0-9])(?=.*[!@#$&=]).{8,}$"
	specialLetters = "!@#$&="
)

/*

(?=.*[A-Z].*[A-Z])        At least 2 upper case letters.
(?=.*[a-z].*[a-z].*[a-z]) At least 3 lower case letters.
(?=.*[0-9].*[0-9])        At least 2 digits.
(?=.*[!@#$&=])            At least 1 special letter.
.{8,}                     At least 8 characters in length.

 */

// var passwordRegexp = regexp.MustCompile(passwordExp) // Golang unsupported Perl syntax: `(?=`

type presents struct {
	uppers, lowers, digits, specials int
}

func validatePresents(p *presents) bool {
	return (p.uppers >= 2) && (p.lowers >= 3) && (p.digits >= 2) && (p.specials >= 1)
}

func IsPasswordValid(password string) bool {
	p := presents {
	}

	if len(password) < 8 {
		return false
	}

	for _, r := range password {
		switch {
			case unicode.IsNumber(r):
				p.digits++

			case unicode.IsUpper(r):
				p.uppers++

			case unicode.IsLower(r):
				p.lowers++

			case strings.IndexRune(specialLetters, r) >= 0:
				p.specials++

			default:
				return false
		}
	}

	return validatePresents(&p)
}
