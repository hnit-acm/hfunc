package basich

type Uint8 uint8

func (i Uint8) GetNative() uint8 {
	return uint8(i)
}

func (i Uint8) GetFunc() Uint8Func {
	return func() uint8 {
		return i.GetNative()
	}
}

type Uint8Func func() uint8

func (g Uint8Func) GetNative() uint8 {
	return g()
}

func (g Uint8Func) ToString(f Uint8ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}
func (g Uint8Func) ToNativeString() string {
	return Uint64(g.GetNative()).GetFunc().ToNativeString()
}
func (g Uint8Func) ToBinaryString() string {
	return Uint64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Uint8Func) ToHexStringLower() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Uint8Func) ToHexStringUpper() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Uint8ToStringFunc func(n uint8) string
