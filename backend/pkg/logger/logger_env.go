package logger

import (
	"os"
	"strings"
)

// GetLogLevelFromEnv returns a log level from an environment variable
func GetLogLevelFromEnv(varName string, defautlLevel LogLevel) LogLevel {
	levelStr := os.Getenv(varName)
	switch strings.ToUpper(levelStr) {
	case "ERROR":
		return LevelError
	case "WARN":
		return LevelWarn
	case "INFO":
		return LevelInfo
	case "DEBUG":
		return LevelDebug
	default:
		return defautlLevel
	}
}
