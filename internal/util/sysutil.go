package util

import "time"

func DateFormat() string {
	return time.Now().Format("20060102150405")
}
