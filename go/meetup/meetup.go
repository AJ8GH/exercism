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
		for i := 1; i <= 7; i++ {
			if isDay(wDay, month, year, i) {
				return i
			}
		}
		return 0
	case Second:
		return 0
	case Third:
		return 0
	case Fourth:
		return 0
	case Last:
		return 0
	case Teenth:
		for i := 13; i <= 19; i++ {
			if isDay(wDay, month, year, i) {
				return i
			}
		}
		return 0
	default:
		panic("bad WeekSchedule")
	}
}

func isDay(wd time.Weekday, m time.Month, y int, d int) bool {
	return false
}
