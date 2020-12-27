// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors


// NewErrNotFound
type ErrNotFound OriginalErrorBase

func NewErrNotFound(msg string) *ErrNotFound {
	err := &ErrNotFound{
		Name: "ErrNotFound",
		Message: msg,
	}
	return err
}

func NewErrNotFoundBy(entity, cond, value string) *ErrNotFound {
	return NewErrNotFound(entity + " by " + cond + ": [" + value + "]")
}

func (err *ErrNotFound) Error() string {
	return err.Name + ": " + err.Message
}


// ErrAlreadyExists
type ErrAlreadyExists OriginalErrorBase

func NewErrAlreadyExists(msg string) *ErrAlreadyExists {
	err := &ErrAlreadyExists{
		Name: "ErrAlreadyExists",
		Message: msg,
	}
	return err
}

func NewErrAlreadyExistsFor(name, value string) *ErrAlreadyExists {
	return NewErrAlreadyExists(name + " for [" + value + "]")
}

func (err *ErrAlreadyExists) Error() string {
	return err.Name + ": " + err.Message
}
