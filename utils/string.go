package utils

import (
	"github.com/hnit-acm/go-common/basic"
)

func StringToSnakeCasedString(str string) string {
	return basic.String(str).SnakeCasedString()
}

func JsonStringToMapStringInterface(jsonStr string) (res map[string]interface{}) {
	return basic.JsonString(jsonStr).GetMapStringInterface()
}
