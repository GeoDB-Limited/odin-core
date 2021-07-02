package rpc

import "time"

func TimeToUTCDate(t time.Time) time.Time {
	year, month, day := t.UTC().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func CalculateDays(t1, t2 time.Time) int {
	return int(t2.UTC().Sub(t1.UTC()).Hours())/24 + 1
}

