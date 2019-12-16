package date

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"
const databaseDateTimeFormat = "2006-01-02 15:04:05"

func GetNow() time.Time {
	return time.Now().UTC()
}

func NowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowForDatabase() string {
	return GetNow().Format(databaseDateTimeFormat)
}
