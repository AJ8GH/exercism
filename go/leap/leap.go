// Package leap allows determining leap year
package leap

// IsLeapYear returns true if year is leap
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
