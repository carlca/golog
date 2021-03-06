package main

import (
	"github.com/kataras/golog"
)

func main() {
	// Default Output is `os.Stderr`,
	// but you can change it:
	// golog.SetOutput(os.Stdout)

	// Time Format defaults to: "2006/01/02 15:04"
	// you can change it to something else or disable it with:
	golog.SetTimeFormat("")

	// Level defaults to "info",
	// but you can change it:
	golog.SetLevel("debug")

	golog.Println("This is a raw message, no levels, no colors.")
	golog.Info("This is an info message, with colors (if the output is terminal)")
	golog.Warn("This is a warning message")
	golog.Error("This is an error message")
	golog.Debug("This is a debug message")
}
