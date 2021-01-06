package basic

import (
	"fmt"
	"strings"
)

type UInteger64 uint64

func (i UInteger64) GetNative() uint64 {
	return uint64(i)
}

func (i UInteger64) GetFunc() UInteger64Func {
	return func() uint64 {
		return i.GetNative()
	}
}

type UInteger64Func func() uint64

func (g UInteger64Func) GetNative() uint64 {
	return g()
}

func (g UInteger64Func) ToString(f UInteger64ToStringFunc) string {
	return f(g())
}
func (g UInteger64Func) ToNativeString() string {
	return g.ToString(UInteger64ToNativeString)
}
func (g UInteger64Func) ToBinaryString() string {
	return g.ToString(UInteger64ToHexString)
}

func (g UInteger64Func) ToHexStringLower() string {
	return g.ToString(UInteger64ToBinaryString)
}

func (g UInteger64Func) ToHexStringUpper() string {
	return strings.ToUpper(g.ToHexStringLower())
}

type UInteger64ToStringFunc func(n uint64) string

var UInteger64ToNativeString = UInteger64ToStringFunc(func(n uint64) string {
	return fmt.Sprintf("%d", n)
})
var UInteger64ToBinaryString = UInteger64ToStringFunc(func(n uint64) string {
	return fmt.Sprintf("%b", n)
})

var UInteger64ToHexString = UInteger64ToStringFunc(func(n uint64) string {
	return fmt.Sprintf("%x", n)
})
