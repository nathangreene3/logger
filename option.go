package logger

import (
	"io"
)

// An Option modifies a logger.
type Option func(*Logger)

// SetFormat returns an option that can set the format for a logger.
func SetFormat(f Format) Option {
	return func(lgr *Logger) {
		lgr.SetFormat(f)
	}
}

// SetOutput returns an option that can set the writer for a logger.
func SetOutput(w io.Writer) Option {
	return func(lgr *Logger) {
		lgr.SetOutput(w)
	}
}
