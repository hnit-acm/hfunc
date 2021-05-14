package arrayh

import "strings"

func ToString(n []string, split string) (res string) {
	if len(n) <= 0 {
		return ""
	}
	var str strings.Builder
	str.Grow(len(n) * (2 + len(split)))
	for k := range n {
		str.WriteString(n[k])
		str.WriteString(split)
	}
	return str.String()[:str.Len()-1]
}
