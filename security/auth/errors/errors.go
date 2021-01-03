// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
)

// ErrUsernameNotFound
func NewErrUsernameNotFound() error {
	return errors.New("UsernameNotFound")
}

// ErrBadCredentials (ErrIncorrectPassword)
func NewErrBadCredentials() error {
	return errors.New("BadCredentials")
}

// ErrAccountStatus:
// (ErrAccountExpired, ErrAccountLocked, ErrCredentialExpired, ErrAccountDisabled)

// ErrAccountExpired
func NewErrAccountExpired() error {
	return errors.New("AccountExpired")
}

// ErrAccountLocked
func NewErrAccountLocked() error {
	return errors.New("AccountLocked")
}

// ErrCredentialExpired
func NewErrCredentialExpired() error {
	return errors.New("CredentialExpired")
}

// ErrAccountDisabled
func NewErrAccountDisabled() error {
	return errors.New("AccountDisabled")
}
