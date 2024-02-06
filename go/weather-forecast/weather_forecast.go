// Package weather provides functionalities to fetch and represent weather data based on geographic locations.
package weather

// CurrentCondition stores the current weather condition of a location.
var CurrentCondition string
// CurrentLocation holds the geographic coordinates of the current location.
var CurrentLocation string

// Forecast() returns the weather forecast for a specified location over a given time period.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
