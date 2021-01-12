package basic

type Uint16 uint16

func (i Uint16) GetNative() uint16 {
	return uint16(i)
}

func (i Uint16) GetFunc() Uint16Func {
	return func() uint16 {
		return i.GetNative()
	}
}

type Uint16Func func() uint16

func (g Uint16Func) GetNative() uint16 {
	return g()
}

func (g Uint16Func) ToString(f Uint16ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}
func (g Uint16Func) ToNativeString() string {
	return Uint64(g.GetNative()).GetFunc().ToNativeString()
}
func (g Uint16Func) ToBinaryString() string {
	return Uint64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Uint16Func) ToHexStringLower() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Uint16Func) ToHexStringUpper() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Uint16ToStringFunc func(n uint16) string