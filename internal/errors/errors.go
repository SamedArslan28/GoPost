package errors

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrEmailConflict = errors.New("email already taken")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrBadRequest    = errors.New("bad request")
	ErrInternal      = errors.New("internal error")
	ErrDuplicated    = errors.New("duplicated")
	ErrParseBody     = errors.New("failed to parse body")
)
