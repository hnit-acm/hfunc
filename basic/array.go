package basic

import (
	"fmt"
	"strings"
)

type Array []interface{}

func (s Array) GetNative() []interface{} {
	return s
}

type ConcatFilter func(s *strings.Builder, elem interface{})

func builderPlus(str *strings.Builder, elem interface{}) {
	switch val := elem.(type) {
	case string:
		str.WriteString(val)
	default:
		str.WriteString(fmt.Sprintf("%v", elem))
	}
}

func (s Array) ToString(split string, c ...ConcatFilter) string {
	if len(s) <= 0 {
		return ""
	}
	filted := false
	if len(c) > 0 {
		filted = true
	}
	var str strings.Builder
	str.Grow(len(s) * (2 + len(split)))
	for k := range s {
		if filted {
			c[0](&str, s[k])
		} else {
			builderPlus(&str, s[k])
		}
		if split != "" {
			str.WriteString(split)
		}
	}
	if split != "" {
		return str.String()[:str.Len()-1]
	}
	return str.String()
}
