package logger

import (
	"bytes"
	"fmt"
)

// Levels
const (
	Info Level = iota
	Warn
	Error
	Fatal
	Debug
)

// A Level indicates the significance of a log entry.
type Level int

// String returns a representation of a level.
func (lvl Level) String() string {
	switch lvl {
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	case Debug:
		return "DEBUG"
	default:
		return ""
	}
}

// ParseLevel parses a representation of a level.
func ParseLevel(level string) (Level, error) {
	switch level {
	case Info.String():
		return Info, nil
	case Warn.String():
		return Warn, nil
	case Error.String():
		return Error, nil
	case Fatal.String():
		return Fatal, nil
	case Debug.String():
		return Debug, nil
	default:
		return 0, ErrInvalidLevel
	}
}

// UnmarshalJSON parses JSON-encoded data into a level.
func (lvl *Level) UnmarshalJSON(b []byte) error {
	if lvl == nil {
		return ErrNilPointer
	}

	b = bytes.Trim(b, `"`)

	parsed, err := ParseLevel(string(b))
	if err != nil {
		return fmt.Errorf("ParseLevel: %w", err)
	}

	*lvl = parsed

	return nil
}

// MarshalJSON returns a JSON-encoded representation of a level.
func (lvl Level) MarshalJSON() ([]byte, error) {
	s := lvl.String()
	if s == "" {
		return nil, ErrInvalidLevel
	}

	return []byte(`"` + s + `"`), nil
}
