package main

import (
	"os"

	"github.com/hnit-acm/hfunc/logh"
	"github.com/hnit-acm/hfunc/logh/stdlog"
)

func main() {

	logger, _ := stdlog.NewLogger(stdlog.Writer(os.Stdout))
	log := logh.NewHelper("task", logger)
	log.Info("hello")
	log.Infof("%d:%s", 123, "hello")
}
