package logger

// Levels
const (
	debugLevel Level = "DEBUG"
	errLevel   Level = "ERROR"
	fatalLevel Level = "FATAL"
	infoLevel  Level = "INFO"
	warnLevel  Level = "WARN"
)

// A Level indicates the significance of a log entry.
type Level string
