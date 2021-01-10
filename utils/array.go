package utils

import "github.com/hnit-acm/hfunc/basic"

func ArrayToString(p []interface{}, split string, funcs ...basic.ConcatFilterFunc) string {
	return basic.Array(p).GetFunc().ToString(split, funcs...)
}

func ArrayStringToString(p []string, split string, funcs ...basic.ConcatFilterFunc) string {
	return basic.ArrayString(p).GetFunc().ToString(split, funcs...)
}
