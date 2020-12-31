package utils

import "github.com/hnit-acm/go-common/basic"

func ArrayToString(p []interface{}, split string, funcs ...basic.ConcatFilter) string {
	return basic.Array(p).ToString(split, funcs...)
}

func ArrayStringToString(p []string, split string, funcs ...basic.ConcatFilter) string {
	return basic.ArrayString(p).ToString(split, funcs...)
}
