package hbasic

type Int8 int8

func (i Int8) GetNative() int8 {
	return int8(i)
}

func (i Int8) GetFunc() Int8Func {
	return func() int8 {
		return i.GetNative()
	}
}

type Int8Func func() int8

func (g Int8Func) GetNative() int8 {
	return g()
}

func (g Int8Func) ToString(f Int8ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}

func (g Int8Func) ToNativeString() string {
	return Int64(g.GetNative()).GetFunc().ToNativeString()
}

func (g Int8Func) ToBinaryString() string {
	return Int64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Int8Func) ToHexStringLower() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Int8Func) ToHexStringUpper() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Int8ToStringFunc func(n int8) string
