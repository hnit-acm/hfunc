package main

import (
	"fmt"
	"go-common/basic"
)

func main() {
	//s := basic.TimeString("2020-01-01")
	var s basic.TimeString = "asd"
	fmt.Println(s.ParseDate())
}
