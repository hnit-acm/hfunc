package logh

import "sync"

var ppFree = sync.Pool{
	New: func() interface{} { return new(printer) },
}

// Logger  interface.
type Logger interface {
	Print(v ...interface{})
}

type printer struct {
	log     Logger
	v       []interface{}
	recycle bool
}

type nopLogger struct{}

func (n *nopLogger) Print(kvpair ...interface{}) {}

func newPrinter() *printer {
	return ppFree.Get().(*printer)
}

func (l *printer) Print(v ...interface{}) {
	l.log.Print(append(v, l.v...)...)
	if l.recycle {
		l.free()
	}
}

func (l *printer) free() {
	l.log = nil
	l.v = nil
	ppFree.Put(l)
}

func with(l Logger, free bool, v ...interface{}) Logger {
	p := newPrinter()
	p.log = l
	p.v = v
	p.recycle = free
	return p
}

func With(l Logger, v ...interface{}) Logger {
	return with(l, false, v)
}

func Debug(l Logger) Logger {
	return with(l, true, LevelKey, LevelDebug)
}

func Info(l Logger) Logger {
	return with(l, true, LevelKey, LevelInfo)
}

func Warn(l Logger) Logger {
	return with(l, true, LevelKey, LevelWarn)
}

func Error(l Logger) Logger {
	return with(l, true, LevelKey, LevelError)
}
