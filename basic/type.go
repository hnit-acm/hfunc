package basic

import (
	"encoding/json"
	"time"
)

type JsonString string

func (j *JsonString) GetNative() string {
	return string(*j)
}

func (j JsonString) GetMapStringInterface() (res map[string]interface{}) {
	err := json.Unmarshal([]byte(j), &res)
	if err != nil {
		return nil
	}
	return
}

type TimeString string

func (t *TimeString) GetNative() string {
	return string(*t)
}

func (t *TimeString) ParseTime() (res *time.Time) {
	ti, err := time.Parse("2006-01-02 15:04:05", t.GetNative())
	if err != nil {
		return nil
	}
	res = &ti
	return
}

func (t *TimeString) ParseDate() (res *time.Time) {
	ti, err := time.Parse("2006-01-02", t.GetNative())
	if err != nil {
		return nil
	}
	res = &ti
	return
}
