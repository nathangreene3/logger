package logger

// A Format defines how log entries are structured.
type Format int

const (
	// Default formats log entries as
	//
	// 	2006-01-02T15:04:05.999999999Z07:00 INFO: Hello, World!
	Default Format = iota

	// JSON formats log entries as
	//
	// 	{
	// 	  "time": "2006-01-02T15:04:05.999999999Z07:00",
	// 	  "level": "INFO",
	// 	  "message": "Hello, World!"
	// 	}
	JSON
)
