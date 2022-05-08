package dates_utils

import "time"

const (
	apiDateLayout = "2006-01-02T23:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}
