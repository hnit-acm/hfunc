package hbasic

import (
	"strings"
)

type ArrayString []string

func (s ArrayString) GetNative() []string {
	return s
}
func (s ArrayString) GetFunc() ArrayStringFunc {
	return func() []string {
		return s.GetNative()
	}
}

type ArrayStringFunc func() []string

func (s ArrayStringFunc) GetNative() []string {
	return s()
}

func (s ArrayStringFunc) ToString(split string, c ...ConcatFilterFunc) (res string) {
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
