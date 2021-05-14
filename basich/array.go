package basich

import (
	"fmt"
	"strings"
)

type Array []interface{}

func (s Array) GetNative() []interface{} {
	return s
}

func (s Array) GetFunc() ArrayFunc {
	return func() []interface{} {
		return s.GetNative()
	}
}

type ArrayFunc func() []interface{}

func (s ArrayFunc) GetNative() []interface{} {
	return s()
}

type ConcatFilterFunc func(s *strings.Builder, elem interface{}) bool

func builderPlus() ConcatFilterFunc {
	return func(str *strings.Builder, elem interface{}) bool {
		switch val := elem.(type) {
		case string:
			str.WriteString(val)
			return true
		default:
			str.WriteString(fmt.Sprintf("%v", elem))
			return true
		}
	}
}

func (s ArrayFunc) ToString(split string, c ...ConcatFilterFunc) string {
	if len(s()) <= 0 {
		return ""
	}
	var str strings.Builder
	str.Grow(len(s()) * (2 + len(split)))
for1:
	for k := range s() {
		for _, filterFunc := range c {
			ok := filterFunc(&str, s()[k])
			if ok {
				if split != "" {
					str.WriteString(split)
				}
				continue for1
			}
		}
		builderPlus()(&str, s()[k])
		if split != "" {
			str.WriteString(split)
		}
	}
	if split != "" {
		return str.String()[:str.Len()-1]
	}
	return str.String()
}
