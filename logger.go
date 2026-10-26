package main

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup logger to be used in these application.
func setupLogger(level zerolog.Level) {
	// Configure zerolog for console output
	cw := zerolog.ConsoleWriter{
		Out:         os.Stdout,
		TimeFormat:  "2006-01-02 15:04:05.000",
		FormatLevel: colorFormatter(),
		FormatCaller: func(i any) string {
			pathStr, ok := i.(string)
			if !ok {
				return "[unknown]"
			}

			return "[" + filepath.Base(pathStr) + "]"
		},
	}

	// Set up global logger
	mw := zerolog.MultiLevelWriter(cw)
	log.Logger = zerolog.New(mw).With().Caller().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)

	// Log something
	log.Debug().Msg("Logger setup completed.")
}

// Create a new formatter for colorized console output.
func colorFormatter() zerolog.Formatter {
	return func(i any) string {
		level, ok := i.(string)
		if !ok {
			return "[UNKNOWN]"
		}

		switch level {
		case "trace":
			return "\x1b[90m[TRACE]\x1b[0m"
		case "debug":
			return "\x1b[32m[DEBUG]\x1b[0m"
		case "info":
			return "\x1b[34m[INFO ]\x1b[0m"
		case "warn":
			return "\x1b[33m[WARN ]\x1b[0m"
		case "error":
			return "\x1b[31m[ERROR]\x1b[0m"
		case "panic":
			return "\x1b[41m[PANIC]\x1b[0m"

		default:
			return "[" + level + "]"
		}
	}
}
