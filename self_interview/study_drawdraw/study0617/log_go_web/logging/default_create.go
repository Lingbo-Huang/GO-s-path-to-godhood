package logging

import (
	"log"
	"os"
)

func NewDefaultLogger(level LogLevel) *DefaultLogger {
	return &DefaultLogger{
		minLevel:     level,
		loggers:      map[LogLevel]*log.Logger{level: log.New(os.Stdout, "", log.LstdFlags)},
		triggerPanic: false,
	}
}
