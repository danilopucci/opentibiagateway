package logger

import (
	"fmt"
	"log/slog"
	"sync"
)

// TestLogger captures logs in memory for testing purposes.
type TestLogger struct {
	mu   sync.Mutex
	Logs []string
}

// NewTestLogger creates a new test logger.
func NewTestLogger() *TestLogger {
	return &TestLogger{}
}

func (l *TestLogger) log(level string, msg string, attrs ...slog.Attr) {
	l.mu.Lock()
	defer l.mu.Unlock()

	full := level + ": " + msg
	for _, attr := range attrs {
		full += fmt.Sprintf(" %s=%v", attr.Key, attr.Value.Any())
	}
	l.Logs = append(l.Logs, full)
}

func (l *TestLogger) Debug(msg string, attrs ...slog.Attr) {
	l.log("DEBUG", msg, attrs...)
}

func (l *TestLogger) Info(msg string, attrs ...slog.Attr) {
	l.log("INFO", msg, attrs...)
}

func (l *TestLogger) Warn(msg string, attrs ...slog.Attr) {
	l.log("WARN", msg, attrs...)
}

func (l *TestLogger) Error(msg string, attrs ...slog.Attr) {
	l.log("ERROR", msg, attrs...)
}

func (l *TestLogger) Debugw(msg string, err error, attrs ...slog.Attr) {
	l.Debug(msg, append(attrs, slog.Any("error", err))...)
}

func (l *TestLogger) Infow(msg string, err error, attrs ...slog.Attr) {
	l.Info(msg, append(attrs, slog.Any("error", err))...)
}

func (l *TestLogger) Warnw(msg string, err error, attrs ...slog.Attr) {
	l.Warn(msg, append(attrs, slog.Any("error", err))...)
}

func (l *TestLogger) Errorw(msg string, err error, attrs ...slog.Attr) {
	l.Error(msg, append(attrs, slog.Any("error", err))...)
}

func (l *TestLogger) With(attrs ...slog.Attr) Logger {
	// For simplicity, don't implement attribute grouping in test logger
	return l
}

// Reset clears captured logs (useful in unit tests)
func (l *TestLogger) Reset() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.Logs = nil
}
