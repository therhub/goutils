package stringutil

import (
	"time"
)

const (
	CONST_SYS_TIME string = "2020-01-01 00:00:00"
)

// system base init time
var systemTime time.Time

func init() {
	systemTime, _ = time.ParseInLocation("2006-01-02 15:04:05", CONST_SYS_TIME, time.Local)
}

func NewID() int64 {

	// get current time
	t := time.Now()

	// clac def time  between time.now and 2020-01-01 00:00:00
	return t.Sub(systemTime).Microseconds()
}
