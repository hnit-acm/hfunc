package basic

import "time"

type TimeString String

func (t TimeString) GetNative() string {
	return string(t)
}

func (t TimeString) GetTime() (res *time.Time) {
	res = t.ParseTime()
	if res != nil {
		return
	}
	res = t.ParseDate()
	if res != nil {
		return
	}
	return
}

type TimeFormatFunc func() string

func defaultTimeFormatFunc() TimeFormatFunc {
	return func() string {
		return "2006-01-02 15:04:05"
	}
}

func defaultDateFormatFunc() TimeFormatFunc {
	return func() string {
		return "2006-01-02"
	}
}

func (t TimeString) ParseTimeFormat(format TimeFormatFunc) (res *time.Time) {
	ti, err := time.Parse(format(), t.GetNative())
	if err != nil {
		return nil
	}
	res = &ti
	return
}

func (t TimeString) ParseTime() (res *time.Time) {
	return t.ParseTimeFormat(defaultTimeFormatFunc())
}

func (t TimeString) ParseDate() (res *time.Time) {
	return t.ParseTimeFormat(defaultDateFormatFunc())
}
