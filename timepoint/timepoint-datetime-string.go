package timepoint

import (
	"fmt"
	"time"
)

func BangkokTimeStringHHMMFromUTC(utcTime time.Time) string {
	gmt7 := time.FixedZone("GMT+7", 7*60*60)
	gmt7Time := utcTime.In(gmt7)

	return gmt7Time.Format("15:04")
}

func BangkokTimeStringHHMMSSFromUTC(t time.Time) string {
	newTime := t.Add(time.Duration(7) * time.Hour)
	return newTime.Format("15:04:05")
}

func GetDisplayDateInThaiFromUTC(utcTime time.Time) string {
	bkkTime := BangkokTimeFromUTC(utcTime)
	month := GetThaiMonthName(bkkTime.Month())

	return fmt.Sprintf("%d %s %d", bkkTime.Day(), month, bkkTime.Year())
}

func GetShortDisplayDateInEngFromUTC(utcTime time.Time) string {
	localTime := utcTime.Add(7 * time.Hour)
	return localTime.Format("02-Jan-06")
}

func GetFullDisplayDateInEngFromUTC(utcTime time.Time) string {
	localTime := utcTime.Add(7 * time.Hour)
	return localTime.Format("02 January 2006")
}

func GetThaiMonthName(month time.Month) string {
	months := []string{
		"",
		"มกราคม",
		"กุมภาพันธ์",
		"มีนาคม",
		"เมษายน",
		"พฤษภาคม",
		"มิถุนายน",
		"กรกฎาคม",
		"สิงหาคม",
		"กันยายน",
		"ตุลาคม",
		"พฤศจิกายน",
		"ธันวาคม",
	}

	return months[int(month)]
}
