package hbasic

type Int16 int16

func (i Int16) GetNative() int16 {
	return int16(i)
}

func (i Int16) GetFunc() Int16Func {
	return func() int16 {
		return i.GetNative()
	}
}

type Int16Func func() int16

func (g Int16Func) GetNative() int16 {
	return g()
}

func (g Int16Func) ToString(f Int16ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}

func (g Int16Func) ToNativeString() string {
	return Int64(g.GetNative()).GetFunc().ToNativeString()
}

func (g Int16Func) ToBinaryString() string {
	return Int64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Int16Func) ToHexStringLower() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Int16Func) ToHexStringUpper() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Int16ToStringFunc func(n int16) string
