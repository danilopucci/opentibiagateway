package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"
)

// TextColorHandler is a slog.Handler that outputs colorful text logs.
type TextColorHandler struct {
	writer  io.Writer
	options *slog.HandlerOptions
}

// NewTextColorHandler creates a new TextColorHandler.
func NewTextColorHandler(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &TextColorHandler{writer: w, options: opts}
}

func (h *TextColorHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.options.Level.Level()
}

func (h *TextColorHandler) Handle(_ context.Context, record slog.Record) error {
	levelColor := map[slog.Level]string{
		slog.LevelDebug: "\033[36mDEBUG\033[0m", // Cyan
		slog.LevelInfo:  "\033[32mINFO\033[0m",  // Green
		slog.LevelWarn:  "\033[33mWARN\033[0m",  // Yellow
		slog.LevelError: "\033[31mERROR\033[0m", // Red
	}

	levelStr, ok := levelColor[record.Level]
	if !ok {
		levelStr = record.Level.String()
	}

	timestamp := record.Time.Format(time.RFC3339)
	msg := record.Message

	// Print the basic log line
	fmt.Fprintf(h.writer, "[%s] %s: %s", timestamp, levelStr, msg)

	// Print the structured attributes
	record.Attrs(func(attr slog.Attr) bool {
		fmt.Fprintf(h.writer, " %s=%v", attr.Key, attr.Value.Any())
		return true
	})

	fmt.Fprintln(h.writer) // end the log line
	return nil
}

func (h *TextColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *TextColorHandler) WithGroup(name string) slog.Handler {
	return h
}
