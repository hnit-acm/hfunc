package basic

import (
	"fmt"
	"strings"
)

type UInteger8 uint8

func (i UInteger8) GetNative() uint8 {
	return uint8(i)
}

func (i UInteger8) GetFunc() UInteger8Func {
	return func() uint8 {
		return i.GetNative()
	}
}

type UInteger8Func func() uint8

func (g UInteger8Func) GetNative() uint8 {
	return g()
}

func (g UInteger8Func) ToString(f UInteger8ToStringFunc) string {
	return f(g())
}

func (g UInteger8Func) ToNativeString() string {
	return g.ToString(UInteger8ToNativeString)
}

func (g UInteger8Func) ToBinaryString() string {
	return g.ToString(UInteger8ToBinaryString)
}

func (g UInteger8Func) ToHexStringLower() string {
	return g.ToString(UInteger8ToHexString)
}

func (g UInteger8Func) ToHexStringUpper() string {
	return strings.ToUpper(g.ToHexStringLower())
}

type UInteger8ToStringFunc func(n uint8) string

var UInteger8ToBinaryString = UInteger8ToStringFunc(func(n uint8) string {
	return fmt.Sprintf("%b", n)
})

var UInteger8ToHexString = UInteger8ToStringFunc(func(n uint8) string {
	return fmt.Sprintf("%x", n)
})

var UInteger8ToNativeString = UInteger8ToStringFunc(func(n uint8) string {
	return fmt.Sprintf("%d", n)
})