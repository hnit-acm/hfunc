package logh

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var (
	logrusInstance *log.Logger
	once           sync.Once
)

type Option func(*Config)

type Config struct {
}

func NewLogrusLog() *logrus.Logger {
	once.Do(func() {
		logrusInstance = logrus.New()
	})
	return logrusInstance
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)
}

func Error(v ...interface{}) {
	log.Error(v...)
}

func Errorf(f string, v ...interface{}) {
	log.Errorf(f, v...)
}

func Warn(v ...interface{}) {
	log.Warn(v...)
}

func Warnf(f string, v ...interface{}) {
	log.Warnf(f, v...)
}

func Info(v ...interface{}) {
	log.Info(v...)
}

func Infof(f string, v ...interface{}) {
	log.Infof(f, v...)
}
