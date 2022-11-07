package logger

import "io"

// An Option modifies a logger.
type Option func(*Logger)

// SetFormat returns an option that can set the format for a logger.
func SetFormat(n Format) Option {
	return func(lgr *Logger) {
		lgr.SetFormat(n)
	}
}

// SetWriter returns an option that can set the writer for a logger.
func SetWriter(w io.Writer) Option {
	return func(lgr *Logger) {
		lgr.SetWriter(w)
	}
}
