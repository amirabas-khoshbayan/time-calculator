package time_calculator

import (
	"time"
)

func GetTodayAndTomorrowUTC() (currentDay, tomorrow time.Time) {
	currentYear, currentMonth, day := time.Now().Date()
	currentDay = time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, time.Now().Location())
	tomorrow = currentDay.AddDate(0, 0, 1)
	return currentDay, tomorrow
}
func GetFirstAndLastOfMonth() (firstOfMonth, lastOfMonth time.Time) {
	currentYear, currentMonth, _ := time.Now().Date()
	firstOfMonth = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.Now().Location())
	lastOfMonth = firstOfMonth.AddDate(0, 1, -1)
	return firstOfMonth , lastOfMonth
}

func CalculateFirstAndLastWeek() (firstWeek, lastWeek time.Time) {
	currentYear, currentMonth, day := time.Now().Date()
	currentDay := time.Date(currentYear, currentMonth, day, 0, 0, 0, 0, time.Now().Location())
	Weekday := currentDay.Weekday()
	firstWeek = time.Time{}
	lastWeek = time.Time{}
	switch Weekday {
	case time.Monday:
		firstWeek = currentDay.AddDate(0, 0, 0)
		lastWeek = currentDay.AddDate(0, 0, 6)
	case time.Tuesday:
		firstWeek = currentDay.AddDate(0, 0, -1)
		lastWeek = currentDay.AddDate(0, 0, 5)
	case time.Wednesday:
		firstWeek = currentDay.AddDate(0, 0, -2)
		lastWeek = currentDay.AddDate(0, 0, 4)
	case time.Thursday:
		firstWeek = currentDay.AddDate(0, 0, -3)
		lastWeek = currentDay.AddDate(0, 0, 3)
	case time.Friday:
		firstWeek = currentDay.AddDate(0, 0, -4)
		lastWeek = currentDay.AddDate(0, 0, 2)
	case time.Saturday:
		firstWeek = currentDay.AddDate(0, 0, -5)
		lastWeek = currentDay.AddDate(0, 0, 1)
	case time.Sunday:
		firstWeek = currentDay.AddDate(0, 0, -6)
		lastWeek = currentDay.AddDate(0, 0, 0)
	}
	return firstWeek, lastWeek
}
func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 6)
	return
}