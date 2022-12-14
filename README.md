# Logger

```go
go get github.com/nathangreene3/logger
```

## Example

```go
func main() {
	var lgr Logger

	defer func() {
		if r := recover(); r != nil {
			lgr.Info(fmt.Sprintf("Recovered from panic: %v", r))
		}
	}()

	var f, err = os.OpenFile("test.log", os.O_CREATE|os.O_RDWR, os.ModePerm)
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
		SetWriter(f),
	)

	lgr.Info("Hello, World!")
	lgr.Warn("Hello?")
	lgr.Error("Goodbye, cruel World.")
	lgr.Stack()
	lgr.Panic("Stopping the World.")
}
```
