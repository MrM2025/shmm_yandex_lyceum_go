package main

import "fmt"

type LogLevel string

const Error, Info LogLevel = "Error", "Info"

type Logger interface {
	Log()
}

type Log struct {
	Level LogLevel
}

func (l Log) Log(error_text string) {
	if l.Level == Error {
		fmt.Printf("ERROR: %s", error_text)
	} else if l.Level == Info {
		fmt.Printf("INFO: %s", error_text)
	}
}

func main() {
	errorLog := &Log{Level: Error}
	errorLog.Log("This is an error message")
}
