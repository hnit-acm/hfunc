package basic

import (
	"crypto/md5"
	"crypto/sha256"
	"hash"
)

type String string

func (s String) GetNative() string {
	return string(s)
}

func (s String) GetFunc() StringFunc {
	return func() string {
		return s.GetNative()
	}
}

type StringFunc func() string

func (s StringFunc) GetNative() string {
	return s()
}

type StringFormatFunc func(string) string

var SnakeCasedStringFormat = StringFormatFunc(func(s string) string {
	newstr := make([]rune, 0)
	for idx, chr := range s {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= 'A' - 'a'
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
})

type StringEncodeFunc func(str string, hash hash.Hash, sum ...string) string

var StringEncode = StringEncodeFunc(
	func(str string, hash hash.Hash, sum ...string) string {
		hash.Write([]byte(str))
		for _, arg := range sum {
			return string(hash.Sum([]byte(arg)))
		}
		return string(hash.Sum([]byte("")))
	},
)

func (s StringFunc) Format(f StringFormatFunc) string {
	return f(s())
}

func (s StringFunc) SnakeCasedString() string {
	return s.Format(SnakeCasedStringFormat)
}

func (s StringFunc) StringEncodeFunc(f StringEncodeFunc, hash hash.Hash, sum ...string) string {
	if f != nil {
		return f(s(), hash, sum...)
	}
	return ""
}

func (s StringFunc) StringEncode(hash hash.Hash, sum ...string) string {
	return s.StringEncodeFunc(StringEncode, hash, sum...)
}

func (s StringFunc) Md5(sum ...string) string {
	return s.StringEncode(md5.New(), sum...)
}

func (s StringFunc) Sha256(sum ...string) string {
	return s.StringEncode(sha256.New(), sum...)
}
