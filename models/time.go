package models

import (
	"strings"
	"time"
)

type Time struct {
	Time     time.Time `json:"time,omitempty"`
	TimeUnix int64     `json:"time_unix,omitempty,string"`
	TimeZone string    `json:"time_zone,omitempty"`
}

func NewTimeNow() Time {
	t := time.Now()
	return Time{Time: t, TimeUnix: t.UnixNano()}
}

func NewTimeWithTime(t time.Time) Time {
	return Time{Time: t, TimeUnix: t.UnixNano()}
}

func NewTimeWithTimeZone(t time.Time, timeZoneString string) Time {
	if timeZoneString == "" {
		timeZoneString = "Local"
	}
	timezone, _ := time.LoadLocation(timeZoneString)
	timeInTimeZone := t.In(timezone)
	return Time{Time: timeInTimeZone, TimeUnix: timeInTimeZone.UnixNano(), TimeZone: timezone.String()}
}

func (t *Time) Day() string {
	return t.Time.Format("2")
}

func (t *Time) Month() string {
	return strings.ToUpper(t.Time.Format("Jan"))
}

func (t *Time) Year() string {
	return t.Time.Format("2006")
}

func (t *Time) FormattedDateNoYear() string {
	return t.Time.Format("Monday, January 2")
}

func (t *Time) FormattedReadableDateOnly() string {
	return t.Time.Format("Monday January 2, 2006 (MST)")
}

func (t *Time) FormattedDateOnly() string {
	return t.Time.Format("2006-01-02")
}

func (t *Time) FormattedTimeOnly() string {
	return t.Time.Format("3:04PM")
}

func (t *Time) FormattedString() string {
	return t.Time.Format("Monday January 2, 2006 15:04")
}
func (t *Time) Short() string {
	return t.Time.Format("Mon 02 Jan 15:04")
}
