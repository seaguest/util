package util

import (
	"fmt"
	"strconv"
	"time"
)

// format time to 2006-01-02 15:04:05 format
func FormatTime(ts interface{}) string {
	t := ParseTime(ts)
	return t.Format("2006-01-02 15:04:05")
}

// parse time from any time format
func ParseTime(timeStr interface{}) (t time.Time) {
	switch v := timeStr.(type) {
	case int, int64, uint64:
		val, _ := strconv.Atoi(fmt.Sprint(v))
		t = time.Unix(int64(val), 0)
	case string:
		t = parse(v)
	case time.Time:
		t = v
	}
	return
}

// parse all kind of time format
func parse(ts string) (t time.Time) {
	t, _ = time.ParseInLocation("2006-01-02", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-1-02", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-1-2 15:04", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-01-02T15:04:05Z", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-01-02T15:04:05", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("2006-01-02T15:04:05Z07:00", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("01/02/2006", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation("20060102", ts, time.Local)
	if !t.IsZero() {
		return
	}

	t, _ = time.ParseInLocation(time.RFC822, ts, time.Local)
	if !t.IsZero() {
		return
	}
	return
}
