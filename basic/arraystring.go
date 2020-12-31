package basic

import (
	"strings"
)

type ArrayString []string

func (s ArrayString) GetNative() []string {
	return s
}

func (s ArrayString) ToString(split string) (res string) {
	if len(s) <= 0 {
		return ""
	}
	var str strings.Builder
	str.Grow(len(s) * (2 + len(split)))
	for k := range s {
		str.WriteString((s)[k])
		if split != "" {
			str.WriteString(split)
		}
	}
	if split != "" {
		return str.String()[:str.Len()-1]
	}
	return str.String()
}
