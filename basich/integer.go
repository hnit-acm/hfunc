package basich

type Int int

func (i Int) GetNative() int {
	return int(i)
}

func (i Int) GetFunc() IntFunc {
	return func() int {
		return i.GetNative()
	}
}

type IntFunc func() int

func (g IntFunc) GetNative() int {
	return g()
}

func (g IntFunc) ToString(f IntToStringFunc) string {
	if f != nil {
		return f(g())
	}
	return g.ToNativeString()
}

func (g IntFunc) ToNativeString() string {
	return Int64(g.GetNative()).GetFunc().ToNativeString()
}

func (g IntFunc) ToBinaryString() string {
	return Int64(g.GetNative()).GetFunc().ToBinaryString()
}

func (g IntFunc) ToHexStringLower() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringLower()
}

func (g IntFunc) ToHexStringUpper() string {
	return Int64(g.GetNative()).GetFunc().ToHexStringUpper()
}

type IntToStringFunc func(n int) string
