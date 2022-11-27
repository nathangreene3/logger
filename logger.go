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
	output io.Writer
	format Format
}

// ---------------------------------------------------------------------
// 	Logging constructors and modifiers
// ---------------------------------------------------------------------

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

// SetFormat sets the format of log entries.
func (lgr *Logger) SetFormat(n Format) {
	lgr.Lock()
	lgr.format = n
	lgr.Unlock()
}

// SetOutput sets the underlying writer.
func (lgr *Logger) SetOutput(w io.Writer) {
	lgr.Lock()
	lgr.output = w
	lgr.Unlock()
}

// ---------------------------------------------------------------------
// 	Logging methods
// ---------------------------------------------------------------------

// Error writes an error message.
func (lgr *Logger) Error(message string) {
	lgr.WriteLogEntry(errLevel, message)
}

// Fatal writes a fatal message, then exits with code 1.
func (lgr *Logger) Fatal(message string) {
	lgr.WriteLogEntry(fatalLevel, message)
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
	lgr.WriteLogEntry(infoLevel, message)
}

// Output returns the underlying writer.
func (lgr *Logger) Output() io.Writer {
	lgr.RLock()
	defer lgr.RUnlock()
	return lgr.output
}

// Panic writes an error, then calls panic.
func (lgr *Logger) Panic(message string) {
	lgr.Error(message)
	panic(message)
}

// Stack writes the current stack as a debug-level message.
func (lgr *Logger) Stack() {
	lgr.WriteLogEntry(debugLevel, string(debug.Stack()))
}

// Warn writes a warning message.
func (lgr *Logger) Warn(message string) {
	lgr.WriteLogEntry(warnLevel, message)
}

// WriteLogEntry writes a log entry to the underlying writer.
func (lgr *Logger) WriteLogEntry(level Level, message string) {
	// defaultFmt corresponds to the default format.
	const defaultFmt = "%s %s: %s\n"

	// logEntry corresponds to the JSON format.
	type logEntry struct {
		Time    time.Time `json:"time"`
		Level   Level     `json:"level"`
		Message string    `json:"message"`
	}

	lgr.Lock()
	defer lgr.Unlock()

	switch lgr.format {
	case Default:
		// Intentionally ignore any error.
		fmt.Fprintf(lgr.output, defaultFmt, time.Now().Format(time.RFC3339Nano), level, message)
	case JSON:
		// Intentionally ignore any error.
		json.NewEncoder(lgr.output).Encode(logEntry{
			Time:    time.Now(),
			Level:   level,
			Message: message,
		})
	default:
		panic("invalid level")
	}
}
