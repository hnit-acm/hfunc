package basich

import "encoding/json"

type JsonString String

func (j JsonString) GetNative() string {
	return string(j)
}

func (j JsonString) GetFunc() JsonStringFunc {
	return func() string {
		return j.GetNative()
	}
}

type JsonStringFunc func() string

func (j JsonStringFunc) GetNative() string {
	return j()
}

func (j JsonStringFunc) GetMapStringInterface() (res map[string]interface{}) {
	err := json.Unmarshal([]byte(j()), &res)
	if err != nil {
		return nil
	}
	return
}
