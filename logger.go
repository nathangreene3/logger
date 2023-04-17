package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// A Logger writes formatted messages to a writer.
type Logger struct {
	sync.RWMutex

	// output is where log entries will be written to. If left empty, log
	// entries will be written to stderr.
	output io.Writer

	// format specifies how log entries will be formatted. If left empty,
	// log entries will be formatted with the default format.
	format Format
}

// New returns a logger. If no options are passed, the logger will write
// to stderr with the default format.
func New(opts ...Option) *Logger {
	var lgr Logger
	lgr.Init(opts...)
	return &lgr
}

// Init modifies a logger. If the underlying writer is unset, it will be
// set to stderr.
func (lgr *Logger) Init(opts ...Option) {
	for _, opt := range opts {
		opt(lgr)
	}

	lgr.Lock()
	if lgr.output == nil {
		lgr.output = os.Stderr
	}

	lgr.Unlock()
}

// Error writes an error message.
func (lgr *Logger) Error(message string) {
	lgr.WriteLogEntry(Error, message)
}

// Fatal writes a fatal-level message, then exits with code 1.
func (lgr *Logger) Fatal(message string) {
	lgr.WriteLogEntry(Fatal, message)
	os.Exit(1)
}

// Format returns the logger's format.
func (lgr *Logger) Format() Format {
	lgr.RLock()
	defer lgr.RUnlock()
	return lgr.format
}

// Info writes an info-level message.
func (lgr *Logger) Info(message string) {
	lgr.WriteLogEntry(Info, message)
}

// Output returns the underlying writer.
func (lgr *Logger) Output() io.Writer {
	lgr.RLock()
	defer lgr.RUnlock()
	return lgr.output
}

// Panic writes an error-level message, then calls panic.
func (lgr *Logger) Panic(message string) {
	lgr.Error(message)
	panic(message)
}

// SetFormat sets the format of log entries.
func (lgr *Logger) SetFormat(f Format) {
	lgr.Lock()
	lgr.format = f
	lgr.Unlock()
}

// SetOutput sets the underlying writer.
func (lgr *Logger) SetOutput(w io.Writer) {
	lgr.Lock()
	lgr.output = w
	lgr.Unlock()
}

// Stack writes the current stack as a debug-level message.
func (lgr *Logger) Stack() {
	lgr.WriteLogEntry(Debug, string(debug.Stack()))
}

// Warn writes a warning message.
func (lgr *Logger) Warn(message string) {
	lgr.WriteLogEntry(Warn, message)
}

// WriteLogEntry writes a log entry to the underlying writer.
func (lgr *Logger) WriteLogEntry(level Level, message string) {
	lgr.Lock()
	defer lgr.Unlock()

	e := LogEntry{
		Time:    time.Now(),
		Level:   level,
		Message: message,
	}

	switch lgr.format {
	case Line:
		if _, err := fmt.Fprintln(lgr.output, e); err != nil {
			panic(fmt.Errorf("fmt.Fprintln: %w", err))
		}
	case JSON:
		if err := json.NewEncoder(lgr.output).Encode(e); err != nil {
			panic(fmt.Errorf("json.Encoder.Encode: %w", err))
		}
	default:
		panic(ErrInvalidFormat)
	}
}
