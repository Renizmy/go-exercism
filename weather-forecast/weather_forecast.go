// Package weather provides tools to determiner forecast weather.
package weather

// CurrentCondition is the current condition.
var CurrentCondition string

// CurrentLocation is the current Location.
var CurrentLocation string

// Forecast returns the current weather condition based on the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
