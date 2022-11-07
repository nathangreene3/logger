package logger

// A Format defines how log entries are formed.
type Format int

const (
	// LevelMessage formats log entries as
	// 	01/02/2006 15:04:05 INFO: Hello, World!
	LevelMessage Format = iota

	// JSON formats log entries as
	// 	{"time":"01/02/2006 15:04:05","level":"INFO","message":"Hello, World!"}.
	JSON
)

// Levels
const (
	debugLevel = "DEBUG"
	errLevel   = "ERROR"
	fatalLevel = "FATAL"
	infoLevel  = "INFO"
	warnLevel  = "WARN"
)

// formatTime defines how log entry time stamps are formatted.
const formatTime = "2006/01/02 15:04:05"

// formats is a list of formats corresponding to the integer format
// constants.
//
//	0 (level-message): 2006/01/02 15:04:05 INFO: Hello, World!
//	1 (json):          {"time":"2006/01/02 15:04:05","level":"INFO","message":"Hello, World!"}
var formats = [...]string{
	"%s %s: %s\n",
	`{"time":%q,"level":%q,"message":%q}` + "\n",
}
