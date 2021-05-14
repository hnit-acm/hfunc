package hbasic

type Uint uint

func (i Uint) GetNative() uint {
	return uint(i)
}

func (i Uint) GetFunc() UintFunc {
	return func() uint {
		return i.GetNative()
	}
}

type UintFunc func() uint

func (g UintFunc) GetNative() uint {
	return g()
}

func (g UintFunc) ToString(f UintToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}
func (g UintFunc) ToNativeString() string {
	return Uint64(g.GetNative()).GetFunc().ToNativeString()
}
func (g UintFunc) ToBinaryString() string {
	return Uint64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g UintFunc) ToHexStringLower() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g UintFunc) ToHexStringUpper() string {
	return Uint64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type UintToStringFunc func(n uint) string
