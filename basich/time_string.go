package basich

import "time"

type TimeString String

func (t TimeString) GetNative() string {
	return string(t)
}

func (t TimeString) GetFunc() TimeStringFunc {
	return func() string {
		return t.GetNative()
	}
}

type TimeStringFunc func() string

func (s TimeStringFunc) GetNative() string {
	return s()
}

func (t TimeStringFunc) TimeFormatFunc(funcs TimeFormatFunc, layout string) (res *time.Time) {
	if funcs != nil {
		return funcs(t(), layout)
	}
	return nil
}

func (t TimeStringFunc) ParseTimeFormat(layout string) (res *time.Time) {
	return t.TimeFormatFunc(TimeFormat, layout)
}

func (t TimeStringFunc) TryGetTime(funcs TimeFormatFunc, layouts ...string) (res *time.Time) {
	if funcs != nil {
		for _, layout := range layouts {
			res = t.TimeFormatFunc(funcs, layout)
			if res != nil {
				return
			}
		}
	}

	for _, layout := range layouts {
		res = t.TimeFormatFunc(TimeFormat, layout)
		if res != nil {
			return
		}
	}

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

type TimeFormatFunc func(timeStr, layout string) *time.Time

var TimeFormat = TimeFormatFunc(
	func(timeStr, layout string) *time.Time {
		ti, err := time.Parse(layout, timeStr)
		if err != nil {
			return nil
		}
		return &ti
	},
)

var timeLayout = func() string {
	return "2006-01-02"
}

var dateLayout = func() string {
	return "2006-01-02"
}

func (t TimeStringFunc) ParseTime() (res *time.Time) {
	return t.ParseTimeFormat(timeLayout())
}

func (t TimeStringFunc) ParseDate() (res *time.Time) {
	return t.ParseTimeFormat(dateLayout())
}
