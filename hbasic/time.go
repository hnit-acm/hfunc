package hbasic

import "time"

type Time time.Time

func (t Time) GetNative() time.Time {
	return time.Time(t)
}

func (t Time) GetFunc() TimeFunc {
	return func() time.Time {
		return time.Time(t)
	}
}

type TimeFormatStringFunc func(t time.Time, layout string) string

var TimeFormatString = TimeFormatStringFunc(func(t time.Time, layout string) string {
	return t.Format(layout)
})

type TimeFunc func() time.Time

func (t TimeFunc) GetNative() time.Time {
	return t()
}

func (t TimeFunc) FormatString(layout string) string {
	return t.TimeFormatStringFunc(TimeFormatString, layout)
}

func (t TimeFunc) TimeFormatStringFunc(funcs TimeFormatStringFunc, layout string) string {
	if funcs != nil {
		return funcs(t(), layout)
	}
	return ""
}

func (t TimeFunc) TryFormat(funcs TimeFormatStringFunc, layouts ...string) string {
	if funcs != nil {
		for _, layout := range layouts {
			res := t.TimeFormatStringFunc(funcs, layout)
			if res != "" {
				return res
			}
		}
	}

	for _, layout := range layouts {
		res := t.TimeFormatStringFunc(TimeFormatString, layout)
		if res != "" {
			return res
		}
	}

	res := t.FormatTime()
	if res != "" {
		return res
	}
	res = t.FormatDate()
	if res != "" {
		return res
	}
	return ""
}

func (t TimeFunc) FormatTime() string {
	res := t.FormatString(timeLayout())
	return res
}

func (t TimeFunc) FormatDate() string {
	res := t.FormatString(dateLayout())
	return res
}
