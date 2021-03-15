package utils

import "encoding/json"

func SourceToTarget(src interface{}, tar interface{}) error {
	srcData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(srcData, tar)
}
