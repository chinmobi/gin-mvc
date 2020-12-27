// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

// Common base error types
type OriginalErrorBase struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
}

type CausedErrorBase struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
	Cause   error   `json:"-"`
}


// ErrLackOfParameter
type ErrLackOfParameter OriginalErrorBase

func newErrLackOfParameter(params []string) *ErrLackOfParameter {
	var msg string
	for i, cnt := 0, len(params); i < cnt; i++ {
		if i > 0 {
			msg += ", "
		}
		msg += params[i]
	}

	err := &ErrLackOfParameter{
		Name: "ErrLackOfParameter",
		Message: msg,
	}
	return err
}

func NewErrLackOfParameter(params ...string) *ErrLackOfParameter {
	return newErrLackOfParameter(params)
}

func (err *ErrLackOfParameter) Error() string {
	return err.Name + ": " + err.Message
}


// ErrInvalidParameter
type ErrInvalidParameter struct {
	Name        string  `json:"name"`
	ParamName   string  `json:"paramName"`
	ParamValue  string  `json:"paramValue"`
	Message     string  `json:"message"`
}

func NewErrInvalidParameter(name, value, msg string) *ErrInvalidParameter {
	err := &ErrInvalidParameter{
		Name: "ErrInvalidParameter",
		ParamName: name,
		ParamValue: value,
		Message: msg,
	}
	return err
}

func NewErrInvalidParameterOf(name, value string) *ErrInvalidParameter {
	return NewErrInvalidParameter(name, value, name + ": [" + value + "]")
}

func (err *ErrInvalidParameter) Error() string {
	return err.Name + ": " + err.Message
}


// ErrInternalError
type ErrInternalError CausedErrorBase

func NewErrInternalError(cause error) *ErrInternalError {
	err := &ErrInternalError{
		Name: "ErrInternalError",
		Message: "Caused by: " + cause.Error(),
		Cause: cause,
	}
	return err
}

func (err *ErrInternalError) Error() string {
	return err.Name + ": " + err.Message
}


// ErrBadRequest
type ErrBadRequest CausedErrorBase

func NewErrBadRequest(cause error) *ErrBadRequest {
	err := &ErrBadRequest{
		Name: "ErrBadRequest",
		Message: "Caused by: " + cause.Error(),
		Cause: cause,
	}
	return err
}

func (err *ErrBadRequest) Error() string {
	return err.Name + ": " + err.Message
}


// ErrMethodNotAllowed
type ErrMethodNotAllowed CausedErrorBase

func NewErrMethodNotAllowed(cause error) *ErrMethodNotAllowed {
	err := &ErrMethodNotAllowed{
		Name: "ErrMethodNotAllowed",
		Message: "Caused by: " + cause.Error(),
		Cause: cause,
	}
	return err
}

func (err *ErrMethodNotAllowed) Error() string {
	return err.Name + ": " + err.Message
}


// ErrNotImplemented
type ErrNotImplemented OriginalErrorBase

func NewErrNotImplemented(msg string) *ErrNotImplemented {
	err := &ErrNotImplemented{
		Name: "ErrNotImplemented",
		Message: msg,
	}
	return err
}

func (err *ErrNotImplemented) Error() string {
	return err.Name + ": " + err.Message
}
