package utils

import "time"

const (
	DateFormat           = "2006-01-02"
	DateTimeFormat       = "2006-01-02 15:04:05"
	DateTimeFormatZsmart = "02-01-2006 15:04:05"
	DateTimeFormatISO    = "2006-01-02T15:04:05Z"
	DateTimeFormatHMAC   = "Mon, 02 Jan 2006 15:04:05 MST"
	TimeZone             = "Asia/Jakarta"
	DateFormatJNE        = "02-01-2006"
)

func ParseDateToString(date time.Time, format string) string {
	return date.Format(format)
}
