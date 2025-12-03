package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

var logger Logger
var currentOpts LoggerOptions

type slogLogger struct {
	logger *slog.Logger
}

// LogLevel represents the minimum log level
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

// String returns the string representation of the log level
func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	default:
		return "info"
	}
}

// ParseLogLevel converts a string to a LogLevel
func ParseLogLevel(s string) LogLevel {
	switch s {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn":
		return LevelWarn
	case "error":
		return LevelError
	default:
		return LevelInfo
	}
}

// LoggerOptions configures the logger behavior
type LoggerOptions struct {
	Level      LogLevel
	AddSource  bool // Add source file/line information
	JSONFormat bool // Use JSON format instead of text
	Writer     io.Writer
}

// DefaultOptions returns sensible default options
func DefaultOptions() LoggerOptions {
	return LoggerOptions{
		Level:      LevelInfo,
		AddSource:  true,
		JSONFormat: false,
		Writer:     os.Stdout,
	}
}

func newSlogLogger(opts LoggerOptions) *slogLogger {
	var level slog.Level
	switch opts.Level {
	case LevelDebug:
		level = slog.LevelDebug
	case LevelInfo:
		level = slog.LevelInfo
	case LevelWarn:
		level = slog.LevelWarn
	case LevelError:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handlerOpts := &slog.HandlerOptions{
		Level:     level,
		AddSource: opts.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Replace source attribute to show only filename
			if a.Key == slog.SourceKey {
				if source, ok := a.Value.Any().(*slog.Source); ok {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}

	var handler slog.Handler
	if opts.JSONFormat {
		handler = slog.NewJSONHandler(opts.Writer, handlerOpts)
	} else {
		handler = slog.NewTextHandler(opts.Writer, handlerOpts)
	}

	return &slogLogger{
		logger: slog.New(handler),
	}
}

func (l *slogLogger) Debug(msg string) {
	l.logger.Debug(strings.TrimRight(msg, "\n"))
}

func (l *slogLogger) Info(msg string) {
	l.logger.Info(strings.TrimRight(msg, "\n"))
}

func (l *slogLogger) Warn(msg string) {
	l.logger.Warn(strings.TrimRight(msg, "\n"))
}

func (l *slogLogger) Error(msg string) {
	l.logger.Error(strings.TrimRight(msg, "\n"))
}

func (l *slogLogger) Fatal(msg string) {
	l.logger.Error(strings.TrimRight(msg, "\n"))
	os.Exit(1)
}

type nopLogger struct{}

func newNOPLogger() *nopLogger {
	return &nopLogger{}
}

func (l *nopLogger) Debug(msg string) {}

func (l *nopLogger) Info(msg string) {}

func (l *nopLogger) Warn(msg string) {}

func (l *nopLogger) Error(msg string) {}

func (l *nopLogger) Fatal(msg string) {}

// DefaultLogger initializes the logger with default settings (Info level, text format)
func DefaultLogger() {
	currentOpts = DefaultOptions()
	logger = newSlogLogger(currentOpts)
}

// NOPLogger initializes a no-operation logger that discards all logs
func NOPLogger() {
	logger = newNOPLogger()
}

// InitLogger initializes the logger with custom options
func InitLogger(opts LoggerOptions) {
	currentOpts = opts
	logger = newSlogLogger(opts)
}

// SetLogLevel changes the log level of the current logger
func SetLogLevel(level LogLevel) {
	currentOpts.Level = level
	logger = newSlogLogger(currentOpts)
}

func SetLogger(l Logger) {
	logger = l
}

func Debug(v ...any) {
	if logger != nil {
		logger.Debug(fmt.Sprintln(v...))
	}
}

func Info(v ...any) {
	if logger != nil {
		logger.Info(fmt.Sprintln(v...))
	}
}

func Warn(v ...any) {
	if logger != nil {
		logger.Warn(fmt.Sprintln(v...))
	}
}

func Error(v ...any) {
	if logger != nil {
		logger.Error(fmt.Sprintln(v...))
	}
}

func Fatal(v ...any) {
	if logger != nil {
		logger.Fatal(fmt.Sprintln(v...))
	}
}
