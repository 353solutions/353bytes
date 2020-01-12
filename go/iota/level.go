package main

import (
	"fmt"
)

type LogLevel uint8

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	MaxLevel = ErrorLevel
)

func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return fmt.Sprintf("unknown level - %d", l)
}

func main() {
	level := WarningLevel
	fmt.Printf("level is %s (%d)\n", level, level)
}
