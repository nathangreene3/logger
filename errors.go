package logger

import (
	"errors"
)

// Errors
var (
	ErrInvalidFormat     = errors.New("invalid format")
	ErrInvalidLevel      = errors.New("invalid error")
	ErrMalformedLogEntry = errors.New("malformed log entry")
)
