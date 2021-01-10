package utils

import (
	"github.com/hnit-acm/hfunc/basic"
	"time"
)

func StringToSnakeCasedString(str string) string {
	return basic.String(str).GetFunc().SnakeCasedString()
}

func JsonStringToMapStringInterface(jsonStr string) (res map[string]interface{}) {
	return basic.JsonString(jsonStr).GetFunc().GetMapStringInterface()
}

func TryTimeStringToTime(timeStr string, layouts ...string) *time.Time {
	return basic.TimeString(timeStr).GetFunc().TryGetTime(basic.TimeFormat, layouts...)
}

func TryTimeToTimeString(t time.Time, layouts ...string) string {
	return basic.Time(t).GetFunc().TryFormat(basic.TimeFormatString, layouts...)
}
