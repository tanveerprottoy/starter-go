package datetime

import "time"

func Now() time.Time {
	return time.Now()
}

func AddDate(years int, months int, days int) time.Time {
	return time.Now().AddDate(years, months, days)
}

func ToUnix(t time.Time) int64 {
	return t.Unix()
}

func SecondMultiply(m time.Duration) time.Duration {
	return time.Second * m
}
