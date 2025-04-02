package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const TimeFormatDay = "2006/01/02"

func GetNowDate() string {
	now := time.Now()
	str := now.Format(TimeFormat) // 使用固定的日期时间格式
	return str
}

func GetNowDay() string {
	now := time.Now()
	str := now.Format(TimeFormatDay) // 使用固定的日期时间格式
	return str
}

const timezone = "Asia/Shanghai"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	if !time.Time(t).IsZero() {
		b = time.Time(t).AppendFormat(b, TimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
