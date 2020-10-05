package date_utils

import "time"

const (
	timeLayout = "2006-01-02 15:04:05.000Z"
)

func GetNow() time.Time{
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(timeLayout)
}