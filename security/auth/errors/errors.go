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

// ErrAccountNotFound
func NewErrAccountNotFound() error {
	return errors.New("AccountNotFound")
}

// ErrBadCredentials
// ErrIncorrectPassword
func NewErrIncorrectPassword() error {
	return errors.New("IncorrectPassword")
}

// ErrCredentialNotFound
func NewErrCredentialNotFound() error {
	return errors.New("CredentialNotFound")
}

// ErrAccountStatus:
// (ErrAccountExpired, ErrCredentialExpired, ErrAccountDisabled, ErrAccountLocked)

// ErrAccountExpired
func NewErrAccountExpired() error {
	return errors.New("AccountExpired")
}

// ErrCredentialExpired
func NewErrCredentialExpired() error {
	return errors.New("CredentialExpired")
}

// ErrAccountDisabled
func NewErrAccountDisabled() error {
	return errors.New("AccountDisabled")
}

// ErrAccountLocked
func NewErrAccountLocked() error {
	return errors.New("AccountLocked")
}
