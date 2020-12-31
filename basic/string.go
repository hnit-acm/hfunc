package basic

type String string

func (s String) GetNative() string {
	return string(s)
}

func (s String) SnakeCasedString() string {
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
}
