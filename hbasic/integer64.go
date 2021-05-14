package hbasic

import (
	"strconv"
	"strings"
)

type Int64 int64

func (i Int64) GetNative() int64 {
	return int64(i)
}

func (i Int64) GetFunc() Int64Func {
	return func() int64 {
		return i.GetNative()
	}
}

type Int64Func func() int64

func (g Int64Func) GetNative() int64 {
	return g()
}

func (g Int64Func) ToString(f Int64ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}

func (g Int64Func) ToNativeString() string {
	return g.ToString(Int64ToNativeString)
}

func (g Int64Func) ToBinaryString() string {
	return g.ToString(Int64ToBinaryString)
}

func (g Int64Func) ToHexStringLower() string {
	return g.ToString(Int64ToHexString)
}

func (g Int64Func) ToHexStringUpper() string {
	return strings.ToUpper(g.ToHexStringLower())
}

type Int64ToStringFunc func(n int64) string

var Int64ToBinaryString = Int64ToStringFunc(func(n int64) string {
	return strconv.FormatInt(n, 2)
})

var Int64ToHexString = Int64ToStringFunc(func(n int64) string {
	return strconv.FormatInt(n, 16)
})

var Int64ToNativeString = Int64ToStringFunc(func(n int64) string {
	return strconv.FormatInt(n, 10)
})
