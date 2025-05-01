package logger

import (
	"log/slog"
	"os"
)

type LoggerBuilder struct {
	options *slog.HandlerOptions
	format  string
}

// NewLoggerBuilder creates a new builder
func NewLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{
		options: &slog.HandlerOptions{
			Level: slog.LevelInfo,
		},
		format: "text", // Default
	}
}

// WithJSONOutput sets JSON format
func (b *LoggerBuilder) WithJSONOutput() *LoggerBuilder {
	b.format = "json"
	return b
}

// WithTextColorOutput sets text color output
func (b *LoggerBuilder) WithTextColorOutput() *LoggerBuilder {
	b.format = "text-color"
	return b
}

// WithPlainTextOutput sets plain text output
func (b *LoggerBuilder) WithPlainTextOutput() *LoggerBuilder {
	b.format = "text"
	return b
}

// WithLogLevel sets the minimum level
func (b *LoggerBuilder) WithLogLevel(level slog.Level) *LoggerBuilder {
	b.options.Level = level
	return b
}

// Build constructs the logger
func (b *LoggerBuilder) Build() Logger {
	var handler slog.Handler

	switch b.format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, b.options)
	case "text-color":
		handler = NewTextColorHandler(os.Stdout, b.options)
	default:
		handler = slog.NewTextHandler(os.Stdout, b.options)
	}

	s := slog.New(handler)
	return &slogWrapper{slog: s}
}
