package basic

import (
	"strconv"
	"strings"
)

type Uint64 uint64


func (i Uint64) GetNative() uint64 {
	return uint64(i)
}

func (i Uint64) GetFunc() Uint64Func {
	return func() uint64 {
		return i.GetNative()
	}
}

type Uint64Func func() uint64

func (g Uint64Func) GetNative() uint64 {
	return g()
}

func (g Uint64Func) ToString(f Uint64ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}
func (g Uint64Func) ToNativeString() string {
	return g.ToString(Uint64ToNativeString)
}
func (g Uint64Func) ToBinaryString() string {
	return g.ToString(Uint64ToHexString)
}

func (g Uint64Func) ToHexStringLower() string {
	return g.ToString(Uint64ToBinaryString)
}

func (g Uint64Func) ToHexStringUpper() string {
	return strings.ToUpper(g.ToHexStringLower())
}

type Uint64ToStringFunc func(n uint64) string

var Uint64ToNativeString = Uint64ToStringFunc(func(n uint64) string {
	return strconv.FormatUint(n, 10)
})
var Uint64ToBinaryString = Uint64ToStringFunc(func(n uint64) string {
	return strconv.FormatUint(n, 2)
})

var Uint64ToHexString = Uint64ToStringFunc(func(n uint64) string {
	return strconv.FormatUint(n, 16)
})
