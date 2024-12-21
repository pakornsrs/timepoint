package timepoint

import "time"

func IsNotZeroTime(t time.Time) bool {
	return !t.IsZero()
}

func IsGivenTimeInPeriodOf(datetime, startDatetime, endDatetime time.Time) bool {
	return datetime.After(startDatetime) && datetime.Before(endDatetime)
}
