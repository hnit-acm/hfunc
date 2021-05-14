package hutils

import (
	"github.com/hnit-acm/hfunc/hbasic"
	"time"
)

func StringToSnakeCasedString(str string) string {
	return hbasic.String(str).GetFunc().SnakeCasedString()
}

func JsonStringToMapStringInterface(jsonStr string) (res map[string]interface{}) {
	return hbasic.JsonString(jsonStr).GetFunc().GetMapStringInterface()
}

func TryTimeStringToTime(timeStr string, layouts ...string) *time.Time {
	return hbasic.TimeString(timeStr).GetFunc().TryGetTime(hbasic.TimeFormat, layouts...)
}

func TryTimeToTimeString(t time.Time, layouts ...string) string {
	return hbasic.Time(t).GetFunc().TryFormat(hbasic.TimeFormatString, layouts...)
}
