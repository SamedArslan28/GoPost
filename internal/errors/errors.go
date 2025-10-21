package errors

import "errors"

var (
	ErrEmailConflict = errors.New("email already taken")
	ErrParseBody     = errors.New("failed to parse body")
	ErrForbidden     = errors.New("user not authorized for this action")
	ErrNotFound      = errors.New("resource not found")
)
