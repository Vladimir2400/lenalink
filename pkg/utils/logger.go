package utils

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
	level  LogLevel
	logger *log.Logger
}

var (
	// Global logger instance
	globalLogger *Logger
)

// init initializes the global logger
func init() {
	globalLogger = NewLogger(INFO)
}

// NewLogger creates a new logger instance
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		logger: log.New(os.Stdout, "", 0),
	}
}

// SetLevel sets the logger level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(message string, args ...interface{}) {
	if l.level <= DEBUG {
		l.log("DEBUG", message, args...)
	}
}

// Info logs an info message
func (l *Logger) Info(message string, args ...interface{}) {
	if l.level <= INFO {
		l.log("INFO", message, args...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(message string, args ...interface{}) {
	if l.level <= WARN {
		l.log("WARN", message, args...)
	}
}

// Error logs an error message
func (l *Logger) Error(message string, args ...interface{}) {
	if l.level <= ERROR {
		l.log("ERROR", message, args...)
	}
}

// log is the internal logging function
func (l *Logger) log(level string, message string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf(message, args...)
	l.logger.Printf("[%s] %s - %s\n", timestamp, level, logMessage)
}

// Global logger functions

// Debug logs a debug message using the global logger
func Debug(message string, args ...interface{}) {
	globalLogger.Debug(message, args...)
}

// Info logs an info message using the global logger
func Info(message string, args ...interface{}) {
	globalLogger.Info(message, args...)
}

// Warn logs a warning message using the global logger
func Warn(message string, args ...interface{}) {
	globalLogger.Warn(message, args...)
}

// Error logs an error message using the global logger
func Error(message string, args ...interface{}) {
	globalLogger.Error(message, args...)
}

// SetLogLevel sets the global logger level
func SetLogLevel(level LogLevel) {
	globalLogger.SetLevel(level)
}

// GetLogger returns the global logger instance
func GetLogger() *Logger {
	return globalLogger
}
