package basic

import "encoding/json"

type JsonString String

func (j JsonString) GetNative() string {
	return string(j)
}

func (j JsonString) GetMapStringInterface() (res map[string]interface{}) {
	err := json.Unmarshal([]byte(j), &res)
	if err != nil {
		return nil
	}
	return
}
