package logger

import (
	"fmt"
)

const (
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	reset  = "\033[0m"
)

type Level int

const (
	Info Level = iota
	Warning
	Error
)

func Print(level Level, message string) {
	switch level {
	case Info:
		fmt.Println(blue + message + reset)
	case Warning:
		fmt.Println(yellow + message + reset)
	case Error:
		fmt.Println(red + message + reset)
	default:
		fmt.Println(message)
	}
}