package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// LogEntry represents a structured log entry.
type LogEntry struct {
	Time    time.Time `json:"time"`
	Level   Level     `json:"level"`
	Message string    `json:"message"`
}

// ParseLogEntry returns a log entry.
func ParseLogEntry(format Format, logEntry string) (LogEntry, error) {
	switch format {
	case Line:
		fields := strings.SplitN(logEntry, " ", 3)
		if len(fields) != 3 {
			return LogEntry{}, ErrMalformedLogEntry
		}

		t, err := time.Parse(time.RFC3339Nano, fields[0])
		if err != nil {
			return LogEntry{}, fmt.Errorf("time.ParseLogEntry: %w", err)
		}

		v, err := ParseLevel(strings.TrimSuffix(fields[1], ":"))
		if err != nil {
			return LogEntry{}, fmt.Errorf("ParseLevel: %w", err)
		}

		e := LogEntry{
			Time:    t,
			Level:   v,
			Message: fields[2],
		}

		return e, nil
	case JSON:
		var e LogEntry
		if err := json.Unmarshal([]byte(logEntry), &e); err != nil {
			return LogEntry{}, fmt.Errorf("json.Unmarshal: %w", err)
		}

		return e, nil
	default:
		return LogEntry{}, ErrInvalidFormat
	}
}

// String returns a representation of a log entry.
func (e LogEntry) String() string {
	const logEntryFmt = "%s %s: %s"
	return fmt.Sprintf(logEntryFmt, e.Time.Format(time.RFC3339Nano), e.Level, e.Message)
}
