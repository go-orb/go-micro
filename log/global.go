package log

// These functions are global functions copied over from the slog library.

import "golang.org/x/exp/slog"

// Trace calls Logger.Trace on the default logger.
func Trace(msg string, args ...any) {
	slog.Default().LogDepth(0, TraceLevel, msg, args...)
}

// Debug calls Logger.Debug on the default logger.
func Debug(msg string, args ...any) {
	slog.Default().LogDepth(0, DebugLevel, msg, args...)
}

// Info calls Logger.Info on the default logger.
func Info(msg string, args ...any) {
	slog.Default().LogDepth(0, InfoLevel, msg, args...)
}

// Warn calls Logger.Warn on the default logger.
func Warn(msg string, args ...any) {
	slog.Default().LogDepth(0, WarnLevel, msg, args...)
}

// Error calls Logger.Error on the default logger.
func Error(msg string, err error, args ...any) {
	if err != nil {
		// TODO: copy over again from the pkg when copy is avoided
		args = append(args[:len(args):len(args)], slog.Any("err", err))
	}

	slog.Default().LogDepth(0, ErrorLevel, msg, args...)
}

// Log calls Logger.Log on the default logger.
func Log(level slog.Level, msg string, args ...any) {
	slog.Default().LogDepth(0, level, msg, args...)
}

// LogAttrs calls Logger.LogAttrs on the default logger.
func LogAttrs(level slog.Level, msg string, attrs ...slog.Attr) { //nolint:revive
	slog.Default().LogAttrsDepth(0, level, msg, attrs...)
}
