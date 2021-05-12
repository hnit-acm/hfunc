package main

import (
	"log"
	"os"
)

type _log struct {
	l *log.Logger
}

var (
	logh = _log{
		l: log.New(os.Stdout, "[hfunc] ", 0),
	}
)

func (l _log) Info(str string, ps ...interface{}) {
	l.l.Printf(str, ps...)
}

func (l _log) Error(err error) {
	l.l.Println(err)
}
