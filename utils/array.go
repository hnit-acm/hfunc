package utils

import "github.com/hnit-acm/go-common/basic"

func ArrayToString(p []interface{}, split string) string {
	return basic.Array(p).ToString(split)
}
