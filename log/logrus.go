package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type LogrusConfig struct {
	EnableConsole bool
	EnableFile    bool
	Structured    bool
	Level         logrus.Level
	FileLocation  string
}

type LogrusLogger struct {
	Logger *logrus.Logger
}

type LogrusNestedLogger struct {
	Logger *logrus.Entry
}

func NewLogrusLogger(cfg LogrusConfig) *logrus.Logger {
	appLogger := logrus.New()

	var output io.Writer
	switch {
	case cfg.EnableConsole && cfg.EnableFile:
		logFile, err := os.OpenFile(cfg.FileLocation, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(fmt.Errorf("unable to setup log file: %w", err))
		}
		output = io.MultiWriter(os.Stderr, logFile)
	case cfg.EnableConsole:
		output = os.Stderr
	case cfg.EnableFile:
		logFile, err := os.OpenFile(cfg.FileLocation, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(fmt.Errorf("unable to setup log file: %w", err))
		}
		output = logFile
	default:
		output = ioutil.Discard
	}

	appLogger.SetOutput(output)
	appLogger.SetLevel(cfg.Level)

	if cfg.Structured {
		appLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			PrettyPrint:       false,
		})
	} else {
		appLogger.SetFormatter(&prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
			ForceFormatting: true,
		})
	}

	return appLogger
}

func InitLogging() {
	cfg := LogrusConfig{
		EnableConsole: true,
		EnableFile:    true,
		Level:         logrus.TraceLevel,
		Structured:    false,
		FileLocation:  "run.log",
	}

	appLogger := NewLogrusLogger(cfg)
	SetLogger(appLogger)
}
func SetLogger(logger *logrus.Logger) {
	Log = logger
}
