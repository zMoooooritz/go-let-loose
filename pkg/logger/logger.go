package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

var logger Logger

type defaultLogger struct {
	logger *log.Logger
}

func newDefaultLogger() *defaultLogger {
	return &defaultLogger{
		logger: log.New(os.Stdout, "go-let-loose: ", log.LstdFlags),
	}
}

func (l *defaultLogger) Debug(msg string) {
	l.logger.Print("[DEBUG] " + msg)
}

func (l *defaultLogger) Info(msg string) {
	l.logger.Print("[INFO] " + msg)
}

func (l *defaultLogger) Warn(msg string) {
	l.logger.Print("[WARN] " + msg)
}

func (l *defaultLogger) Error(msg string) {
	l.logger.Print("[ERROR] " + msg)
}

func (l *defaultLogger) Fatal(msg string) {
	l.logger.Print("[FATAL] " + msg)
}

type nopLogger struct{}

func newNOPLogger() *nopLogger {
	return &nopLogger{}
}

func (l *nopLogger) Debug(msg string) {
}

func (l *nopLogger) Info(msg string) {
}

func (l *nopLogger) Warn(msg string) {
}

func (l *nopLogger) Error(msg string) {
}

func (l *nopLogger) Fatal(msg string) {
}

func DefaultLogger() {
	logger = newDefaultLogger()
}

func NOPLogger() {
	logger = newNOPLogger()
}

func SetLogger(l Logger) {
	logger = l
}

func Debug(v ...interface{}) {
	if logger != nil {
		logger.Debug(fmt.Sprintln(v...))
	}
}

func Info(v ...interface{}) {
	if logger != nil {
		logger.Info(fmt.Sprintln(v...))
	}
}

func Warn(v ...interface{}) {
	if logger != nil {
		logger.Warn(fmt.Sprintln(v...))
	}
}

func Error(v ...interface{}) {
	if logger != nil {
		logger.Error(fmt.Sprintln(v...))
	}
}

func Fatal(v ...interface{}) {
	if logger != nil {
		logger.Fatal(fmt.Sprintln(v...))
	}
}
