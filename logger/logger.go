package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *Logger) Log(level LogLevel, message string, context interface{}, err error) {
	prefix := ""

	switch level {
	case DEBUG:
		prefix = "[DEBUG] "
	case INFO:
		prefix = "[INFO] "
	case WARNING:
		prefix = "[WARNING] "
	case ERROR:
		prefix = "[ERROR] "
	default:
		prefix = "[UNKNOWN] "
	}

	logMessage := prefix + message

	if err != nil {
		logMessage += " | Error: " + err.Error()
	}

	if context != nil {
		logMessage += " | Context: " + fmt.Sprintf("%+v", context)
	}

	l.logger.Println(logMessage)
}
