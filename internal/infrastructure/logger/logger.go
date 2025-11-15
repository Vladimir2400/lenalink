package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel defines the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger represents the application logger
type Logger struct {
	level LogLevel
	log   *log.Logger
}

// New creates a new logger instance
func New(level string) *Logger {
	var logLevel LogLevel
	switch level {
	case "DEBUG":
		logLevel = DEBUG
	case "WARN":
		logLevel = WARN
	case "ERROR":
		logLevel = ERROR
	default:
		logLevel = INFO
	}

	return &Logger{
		level: logLevel,
		log:   log.New(os.Stdout, "", 0),
	}
}

// Debug logs a debug message
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log.Printf("[%s] DEBUG - %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}

// Info logs an info message
func (l *Logger) Info(format string, args ...interface{}) {
	if l.level <= INFO {
		l.log.Printf("[%s] INFO - %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, args ...interface{}) {
	if l.level <= WARN {
		l.log.Printf("[%s] WARN - %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}

// Error logs an error message
func (l *Logger) Error(format string, args ...interface{}) {
	if l.level <= ERROR {
		l.log.Printf("[%s] ERROR - %s\n", time.Now().Format("2006-01-02 15:04:05"), fmt.Sprintf(format, args...))
	}
}
