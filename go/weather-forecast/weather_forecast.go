// Package weather contains information and functions to
// forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition current condition of weather forecast.
var CurrentCondition string

// CurrentLocation current city of weather forecast.
var CurrentLocation string

// Forecast provides the current weather forecast for a given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
