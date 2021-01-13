package main

import "regexp"

var expNewService, _ = regexp.Compile(`^new \S+$`)
