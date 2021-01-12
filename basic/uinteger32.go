package basic

type Uint32 uint32

func (i Uint32) GetNative() uint32 {
	return uint32(i)
}

func (i Uint32) GetFunc() Uint32Func {
	return func() uint32 {
		return i.GetNative()
	}
}

type Uint32Func func() uint32

func (g Uint32Func) GetNative() uint32 {
	return g()
}

func (g Uint32Func) ToString(f Uint32ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}
func (g Uint32Func) ToNativeString() string {
	return Uint64(g.GetNative()).GetFunc().ToNativeString()
}
func (g Uint32Func) ToBinaryString() string {
	return Uint64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Uint32Func) ToHexStringLower() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Uint32Func) ToHexStringUpper() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Uint32ToStringFunc func(n uint32) string