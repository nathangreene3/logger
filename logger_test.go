package logger

import (
	"fmt"
	"os"
	"testing"
)

func TestJSON(t *testing.T) {
	var lgr Logger
	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		lgr.Fatal(err.Error())
	}

	defer func() {
		if err := f.Close(); err != nil {
			lgr.Error(err.Error())
		}
	}()

	lgr.Init(
		SetFormat(JSON),
		SetOutput(f),
	)

	defer func() {
		if r := recover(); r != nil {
			lgr.Info(fmt.Sprintf("Recovered from panic: %v", r))
			lgr.Stack()
		}
	}()

	lgr.Info("This is an informational message.")
	lgr.Warn("This is a warning message.")
	lgr.Error("This is an error message.")
	lgr.Panic("This is a panic message (an error with panic).")

	// Note: Can't call Logger.Fatal. It causes tests to fail and can't be
	// recovered.
}
