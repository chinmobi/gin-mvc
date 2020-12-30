// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

type CausedErrorBase struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
	Cause   error   `json:"-"`
}

// ErrAuthentication
type ErrAuthentication CausedErrorBase

func NewErrAuthentication(cause error) *ErrAuthentication {
	err := &ErrAuthentication{
		Name: "ErrAuthentication",
		Message: "Caused by: " + cause.Error(),
		Cause: cause,
	}
	return err
}

func (err *ErrAuthentication) Error() string {
	return err.Name + ": " + err.Message
}

// ErrAccessDenied
type ErrAccessDenied CausedErrorBase

func NewErrAccessDenied(cause error) *ErrAccessDenied {
	err := &ErrAccessDenied{
		Name: "ErrAccessDenied",
		Message: "Caused by: " + cause.Error(),
		Cause: cause,
	}
	return err
}

func (err *ErrAccessDenied) Error() string {
	return err.Name + ": " + err.Message
}
