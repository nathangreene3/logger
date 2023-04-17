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
func (v Level) String() string {
	switch v {
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
func ParseLevel(s string) (Level, error) {
	switch s {
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
func (v *Level) UnmarshalJSON(b []byte) error {
	if v == nil {
		return ErrNilPointer
	}

	b = bytes.Trim(b, `"`)

	lvl, err := ParseLevel(string(b))
	if err != nil {
		return fmt.Errorf("ParseLevel: %w", err)
	}

	*v = lvl

	return nil
}

// MarshalJSON returns a JSON-encoded representation of a level.
func (v Level) MarshalJSON() ([]byte, error) {
	s := v.String()
	if s == "" {
		return nil, ErrInvalidLevel
	}

	return []byte(`"` + s + `"`), nil
}
