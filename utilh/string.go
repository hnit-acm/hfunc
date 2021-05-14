package utilh

import (
	"github.com/hnit-acm/hfunc/basich"
	"time"
)

func StringToSnakeCasedString(str string) string {
	return basich.String(str).GetFunc().SnakeCasedString()
}

func JsonStringToMapStringInterface(jsonStr string) (res map[string]interface{}) {
	return basich.JsonString(jsonStr).GetFunc().GetMapStringInterface()
}

func TryTimeStringToTime(timeStr string, layouts ...string) *time.Time {
	return basich.TimeString(timeStr).GetFunc().TryGetTime(basich.TimeFormat, layouts...)
}

func TryTimeToTimeString(t time.Time, layouts ...string) string {
	return basich.Time(t).GetFunc().TryFormat(basich.TimeFormatString, layouts...)
}
