package hbasic

type Int32 int32

func (i Int32) GetNative() int32 {
	return int32(i)
}

func (i Int32) GetFunc() Int32Func {
	return func() int32 {
		return i.GetNative()
	}
}

type Int32Func func() int32

func (g Int32Func) GetNative() int32 {
	return g()
}

func (g Int32Func) ToString(f Int32ToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}

func (g Int32Func) ToNativeString() string {
	return Int64(g.GetNative()).GetFunc().ToNativeString()
}

func (g Int32Func) ToBinaryString() string {
	return Int64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g Int32Func) ToHexStringLower() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g Int32Func) ToHexStringUpper() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type Int32ToStringFunc func(n int32) string
