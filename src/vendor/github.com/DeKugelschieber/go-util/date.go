package util

import (
	"time"
)

const (
	date_format = "2006-01-02 15:04:05"
)

func ParseMySQLTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(date_format, timestamp)
}
