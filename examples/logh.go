package main

import (
	"os"

	"github.com/hnit-acm/hfunc/hlog"
	"github.com/hnit-acm/hfunc/hlog/stdlog"
)

func main() {

	logger, _ := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	log := hlog.NewHelper("task", logger)
	log.Info("hello")
	log.Infof("%d:%s", 123, "hello")
}
