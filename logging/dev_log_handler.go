package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

type devLogHandler struct {
}

func (l *devLogHandler) Enabled(c context.Context, level slog.Level) bool {
	return true
}

func (l *devLogHandler) Handle(c context.Context, record slog.Record) error {
	attrs := map[string]any{}

	line := fmt.Sprintf(
		"%04d-%02d-%02d %02d:%02d:%02d [%s] %s",
		record.Time.Year(),
		record.Time.Month(),
		record.Time.Day(),
		record.Time.Hour(),
		record.Time.Minute(),
		record.Time.Second(),
		strings.ToUpper(record.Level.String()),
		record.Message,
	)

	record.Attrs(func(a slog.Attr) bool {
		attrs[a.Key] = a.Value.Any()
		return true
	})

	if len(attrs) > 0 {
		s, _ := json.MarshalIndent(attrs, "", "  ")
		line += " " + string(s)
	}

	fmt.Println(line)
	return nil
}

func (l *devLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("not implemented")
}

func (l *devLogHandler) WithGroup(name string) slog.Handler {
	panic("not implemented")
}