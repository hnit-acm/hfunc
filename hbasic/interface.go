package hbasic

type BasicTypeString interface {
	GetNative() string
}

type BasicTypeUint8 interface {
	GetNative() uint8
}
type BasicTypeUint16 interface {
	GetNative() uint16
}
type BasicTypeUint32 interface {
	GetNative() uint32
}
type BasicTypeUint64 interface {
	GetNative() uint64
}
type BasicTypeUint interface {
	GetNative() uint
}

type BasicTypeInt8 interface {
	GetNative() int8
}
type BasicTypeInt16 interface {
	GetNative() int16
}
type BasicTypeInt32 interface {
	GetNative() int32
}
type BasicTypeInt64 interface {
	GetNative() int64
}
type BasicTypeInt interface {
	GetNative() int
}

type BasicTypeArrayString interface {
	GetNative() []string
}

type BasicTypeArray interface {
	GetNative() []interface{}
}
