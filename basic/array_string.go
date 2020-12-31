package basic

import (
	"strings"
)

type ArrayString []string

func (s ArrayString) GetNative() []string {
	return s
}

func (s ArrayString) ToString(split string, c ...ConcatFilter) (res string) {
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
