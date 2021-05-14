package hutils

import "github.com/hnit-acm/hfunc/hbasic"

func ArrayToString(p []interface{}, split string, funcs ...hbasic.ConcatFilterFunc) string {
	return hbasic.Array(p).GetFunc().ToString(split, funcs...)
}

func ArrayStringToString(p []string, split string, funcs ...hbasic.ConcatFilterFunc) string {
	return hbasic.ArrayString(p).GetFunc().ToString(split, funcs...)
}
