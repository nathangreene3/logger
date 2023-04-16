package logger

import (
	"fmt"
	"os"
	"testing"
)

func TestJSON(t *testing.T) {
	var lgr Logger

	defer func() {
		if r := recover(); r != nil {
			lgr.Info(fmt.Sprintf("Recovered from panic: %v", r))
		}
	}()

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

	lgr.Info("Hello, World!")
	lgr.Warn("Hello?")
	lgr.Error("Goodbye, cruel World.")
	lgr.Stack()
	lgr.Panic("Stopping the World.")
}
