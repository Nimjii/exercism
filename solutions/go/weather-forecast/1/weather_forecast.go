// Package weather provides weather forecast data.
package weather

var (
    // CurrentCondition represents the current weather condition.
	CurrentCondition string
    // CurrentLocation represents the location where the weather condition was determined.
	CurrentLocation  string
)

// Forecast provides information about the weather condition in a certain location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
