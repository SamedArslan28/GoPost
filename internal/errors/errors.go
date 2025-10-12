package errors

import "errors"

var (
	ErrEmailConflict = errors.New("email already taken")
	ErrParseBody     = errors.New("failed to parse body")
)
