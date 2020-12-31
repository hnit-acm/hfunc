package utils

import (
	"github.com/hnit-acm/go-common/basic"
	"time"
)

func StringToSnakeCasedString(str string) string {
	return basic.String(str).SnakeCasedString()
}

func JsonStringToMapStringInterface(jsonStr string) (res map[string]interface{}) {
	return basic.JsonString(jsonStr).GetMapStringInterface()
}

func TimeStringToTime(timeStr string, funcs ...basic.TimeFormatFunc) *time.Time {
	return basic.TimeString(timeStr).GetTime(funcs...)
}
