package logh

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	logrusInstance *log.Logger
	once           sync.Once
)

type Option func(*Config)

type Config struct {
	level Level
}

type Helper struct {
	cfgs Config
}

func NewLogrusLog() *log.Logger {
	once.Do(func() {
		logrusInstance = log.New()
	})
	return logrusInstance
}

func NewHelper(opts ...Option) *Helper {
	configs := Config{}
	for _, o := range opts {
		o(&configs)
	}
	return &Helper{cfgs: configs}
}

func (h *Helper) Info(v ...interface{}) {
	log.Info(v...)
}

func (h *Helper) Infof(f string, v ...interface{}) {
	log.Infof(f, v...)
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
