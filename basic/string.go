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

type StringFormatFunc func(string) string

var SnakeCasedStringFormat = StringFormatFunc(func(s string) string {
	newstr := make([]rune, 0)
	for idx, chr := range s {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '_')
			}
			chr -= ('A' - 'a')
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

func (s String) SnakeCasedString() string {
	return SnakeCasedStringFormat(string(s))
}

func (s String) Md5(sum ...string) string {
	return s.Encode(md5.New(), sum...)
}

func (s String) Sha256(sum ...string) string {
	return s.Encode(sha256.New(), sum...)
}

func (s String) Encode(hash hash.Hash, sum ...string) string {
	return StringEncode(string(s), hash, sum...)
}
