package main

import (
	"flag"
	"fmt"
	"github.com/hnit-acm/hfunc/basic"
	"github.com/hnit-acm/hfunc/utils"
)

func main() {
	flag.Parse()
	if !flag.Parsed() {
		return
	}
	args := flag.Args()
	fmt.Println(args)
	argsString := utils.ArrayStringToString(args, " ")
	//expNewService, _ := regexp.Compile(`^new \S+$`)
	//expNewService, _ := regexp.Compile(`^new \S+$`)
	switch {
	case expNewService.MatchString(argsString):
		{
			fmt.Println("new service ", args[1])
			newService(basic.String(args[1]))
			return
		}
	}
}
