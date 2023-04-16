package logger

// Levels
const (
	Debug Level = "DEBUG"
	Error Level = "ERROR"
	Fatal Level = "FATAL"
	Info  Level = "INFO"
	Warn  Level = "WARN"
)

// A Level indicates the significance of a log entry.
type Level string
