// Package weather provides a way to get the weather forecast.
package weather

// CurrentCondition Represents the latest condition provided to the forecast.
var CurrentCondition string

// CurrentLocation Represents the latest location provided to the forecast.
var CurrentLocation string

// Forecast returns a string formatted with the condition for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
