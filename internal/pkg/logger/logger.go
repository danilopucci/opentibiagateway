package logger

import (
	"log/slog"
)

type Logger interface {
	Debug(msg string, attrs ...slog.Attr)
	Info(msg string, attrs ...slog.Attr)
	Warn(msg string, attrs ...slog.Attr)
	Error(msg string, attrs ...slog.Attr)

	Debugw(msg string, err error, attrs ...slog.Attr)
	Infow(msg string, err error, attrs ...slog.Attr)
	Warnw(msg string, err error, attrs ...slog.Attr)
	Errorw(msg string, err error, attrs ...slog.Attr)

	With(attrs ...slog.Attr) Logger
}

// Internal wrapper to slog
type slogWrapper struct {
	slog *slog.Logger
}

func (l *slogWrapper) Debug(msg string, attrs ...slog.Attr) {
	l.slog.Debug(msg, toAny(attrs)...)
}

func (l *slogWrapper) Info(msg string, attrs ...slog.Attr) {
	l.slog.Info(msg, toAny(attrs)...)
}

func (l *slogWrapper) Warn(msg string, attrs ...slog.Attr) {
	l.slog.Warn(msg, toAny(attrs)...)
}

func (l *slogWrapper) Error(msg string, attrs ...slog.Attr) {
	l.slog.Error(msg, toAny(attrs)...)
}

func (l *slogWrapper) Debugw(msg string, err error, attrs ...slog.Attr) {
	attrs = append(attrs, slog.Any("error", err))
	l.slog.Debug(msg, toAny(attrs)...)
}

func (l *slogWrapper) Infow(msg string, err error, attrs ...slog.Attr) {
	attrs = append(attrs, slog.Any("error", err))
	l.slog.Info(msg, toAny(attrs)...)
}

func (l *slogWrapper) Warnw(msg string, err error, attrs ...slog.Attr) {
	attrs = append(attrs, slog.Any("error", err))
	l.slog.Warn(msg, toAny(attrs)...)
}

func (l *slogWrapper) Errorw(msg string, err error, attrs ...slog.Attr) {
	attrs = append(attrs, slog.Any("error", err))
	l.slog.Error(msg, toAny(attrs)...)
}

func (l *slogWrapper) With(attrs ...slog.Attr) Logger {
	return &slogWrapper{slog: l.slog.With(toAny(attrs)...)}
}

// Helper to convert []slog.Attr → []any
func toAny(attrs []slog.Attr) []any {
	out := make([]any, 0, len(attrs))
	for _, attr := range attrs {
		out = append(out, attr)
	}
	return out
}
