package basic

type JsonString string

func (j JsonString) GetNative() string {
	return string(j)
}
