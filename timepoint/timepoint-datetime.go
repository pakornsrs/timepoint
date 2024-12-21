package timepoint

import "time"

func UTCTimeNow() time.Time {
	return time.Now().UTC()
}

func BangkokTimeNow() time.Time {
	gmt7 := time.FixedZone("GMT+7", 7*60*60)
	timeUTC := UTCTimeNow()

	return timeUTC.In(gmt7)
}

func BangkokTimeFromUTC(utcTime time.Time) time.Time {
	gmt7 := time.FixedZone("GMT+7", 7*60*60)

	return utcTime.In(gmt7)
}

func NumberOfDaysInMonth(year int, month time.Month) int {
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	nextMonth := firstDay.AddDate(0, 1, 0)

	return int(nextMonth.Sub(firstDay).Hours() / 24)
}

func GetDaySuffix(day int) string {

	if day >= 11 && day <= 13 {
		return "th"
	}

	switch day % 10 {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}
