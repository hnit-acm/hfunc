package basic

type BasicTypeString interface {
	GetNative() string
}

type BasicTypeUint8 interface {
	GetNative() uint8
}

type BasicTypeUint64 interface {
	GetNative() uint64
}

type BasicTypeArrayString interface {
	GetNative() []string
}
