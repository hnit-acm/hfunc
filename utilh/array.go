package utilh

import "github.com/hnit-acm/hfunc/basich"

func ArrayToString(p []interface{}, split string, funcs ...basich.ConcatFilterFunc) string {
	return basich.Array(p).GetFunc().ToString(split, funcs...)
}

func ArrayStringToString(p []string, split string, funcs ...basich.ConcatFilterFunc) string {
	return basich.ArrayString(p).GetFunc().ToString(split, funcs...)
}
