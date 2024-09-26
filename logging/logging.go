package logging

import (
	"log/slog"
	"os"
)

func SetupDevLogger() {
	slog.SetDefault(slog.New(&devLogHandler{}))
}

func SetupProdLogger() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "level" {
				a.Key = "severity"
			}
			if a.Key == "msg" {
				a.Key = "message"
			}
			return a
		},
	})))
}

func SetupLoggerByEnv(variableName string, expected string) {
	if os.Getenv(variableName) == expected {
		SetupProdLogger()
	} else {
		SetupDevLogger()
	}
}
