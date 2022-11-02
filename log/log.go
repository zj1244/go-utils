package log

type Logger interface {
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

var Log Logger

func init() {
	InitLogging()
}
func Error(args ...interface{}) {
	Log.Error(args...)
}
func Errorf(format string, args ...interface{}) {

	Log.Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}
func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}
