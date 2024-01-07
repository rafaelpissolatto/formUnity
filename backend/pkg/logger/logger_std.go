package logger

import "log"

type LogLevel int

const (
	LevelError LogLevel = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

type LoggerStd struct {
	level LogLevel
}

func NewLoggerStd(level LogLevel) *LoggerStd {
	return &LoggerStd{level: level}
}

func (l *LoggerStd) log(level LogLevel, prefix string, v ...interface{}) {
	if l.level >= level {
		v = append([]interface{}{prefix}, v...)
		log.Println(v...)
	}
}

func (l *LoggerStd) Info(v ...interface{}) {
	l.log(LevelInfo, "INFO: ", v)
}

func (l *LoggerStd) Warn(v ...interface{}) {
	l.log(LevelWarn, "WARN: ", v)
}

func (l *LoggerStd) Error(v ...interface{}) {
	l.log(LevelError, "ERROR: ", v)
}

func (l *LoggerStd) Debug(v ...interface{}) {
	l.log(LevelDebug, "DEBUG: ", v)
}
