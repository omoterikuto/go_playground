package main

import (
	"fmt"
	"time"
)

const (
	secondsPerMinute  = 60
	secondsPerHour    = 60 * secondsPerMinute
	jstTimeDifference = 9 * secondsPerHour
)

func main() {
	now := time.Date(2024, 8, 2, 0, 0, 0, 0, JST)
	fmt.Println(LastWeekday(now, time.Monday))
}

var JST = time.FixedZone("JST", jstTimeDifference)

func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.In(JST).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, JST)
}

func LastWeekday(now time.Time, wday time.Weekday) time.Time {
	cur := BeginningOfDay(now)
	for cur.Weekday() != wday {
		cur = cur.AddDate(0, 0, -1)
	}
	return cur
}
