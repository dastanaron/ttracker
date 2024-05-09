package utils

import (
	"fmt"
	"math"
	"time"
)

func ConvertSecondsToHumanTime(seconds int) string {
	hours := math.Floor(float64(seconds) / 3600)
	minutes := math.Floor(float64(seconds%3600) / 60)
	remainingSeconds := seconds % 60

	formedHours := ""

	if hours < 10 {
		formedHours = fmt.Sprintf("0%d", int(hours))
	} else {
		formedHours = fmt.Sprintf("%d", int(hours))
	}

	formedMinutes := ""

	if minutes < 10 {
		formedMinutes = fmt.Sprintf("0%d", int(minutes))
	} else {
		formedMinutes = fmt.Sprintf("%d", int(minutes))
	}

	formedSeconds := ""

	if remainingSeconds < 10 {
		formedSeconds = fmt.Sprintf("0%d", int(remainingSeconds))
	} else {
		formedSeconds = fmt.Sprintf("%d", int(remainingSeconds))
	}

	return fmt.Sprintf("%s:%s:%s", formedHours, formedMinutes, formedSeconds)
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
