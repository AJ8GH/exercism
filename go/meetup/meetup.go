package meetup

import "time"

type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
	Teenth WeekSchedule = "teenth"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	switch wSched {
	case First:
		return findDay(year, month, wDay, 1, 7, 1)
	case Second:
		return findDay(year, month, wDay, 1, 14, 2)
	case Third:
		return findDay(year, month, wDay, 1, 21, 3)
	case Fourth:
		return findDay(year, month, wDay, 1, 28, 4)
	case Last:
		return findLastDay(year, month, wDay)
	case Teenth:
		return findDay(year, month, wDay, 13, 19, 1)
	default:
		panic("bad WeekSchedule")
	}
}

func findDay(
	y int,
	m time.Month,
	wd time.Weekday,
	startDay,
	endDay,
	occurrence int,
) int {
	found := 0
	for i := startDay; i <= endDay; i++ {
		if isDay(wd, m, y, i) {
			found++
			if found == occurrence {
				return i
			}
		}
	}
	return 0
}

func findLastDay(y int, m time.Month, wd time.Weekday) int {
	startDay := 31
	switch m {
	case time.February:
		startDay = daysInFeb(y)
	case time.April, time.June, time.September, time.November:
		startDay = 30
	}
	for i := startDay; i >= startDay-7; i-- {
		if isDay(wd, m, y, i) {
			return i
		}
	}
	return 0
}

func isDay(wd time.Weekday, m time.Month, y int, d int) bool {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Weekday() == wd
}

func daysInFeb(y int) int {
	switch {
	case y%4 == 0 && (y%100 != 0 || y%400 == 0):
		return 29
	default:
		return 28
	}
}
