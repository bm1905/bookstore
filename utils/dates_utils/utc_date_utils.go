package dates_utils

import "time"

const apiDateLayout = "2006-01-02T23:04:05Z"

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
