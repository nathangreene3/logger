package logger

import (
	"errors"
)

// Errors
var (
	ErrInvalidFormat     = errors.New("invalid format")
	ErrInvalidLevel      = errors.New("invalid level")
	ErrMalformedLogEntry = errors.New("malformed log entry")
	ErrNilPointer        = errors.New("nil pointer")
)
