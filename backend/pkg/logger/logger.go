package logger

type Logger interface {
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debug(v ...interface{})
}
