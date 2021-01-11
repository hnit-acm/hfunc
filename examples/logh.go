package main

import (
	"github.com/hnit-acm/hfunc/logh"
)

func main() {
	// logh.Info("hello", "fff")
	// logh.Infof("%d:%s", 123, "hello")
	// logl := logh.NewLogrusLog()
	// logl.Infof("%d:%s", 1234, "helloworld")
	log := logh.NewHelper()
	log.Info("hello")
}
