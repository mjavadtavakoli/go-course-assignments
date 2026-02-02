package main

import "fmt"

func toFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func toCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func toKelvin(c float64) float64 {
	return c + 273.15
}

func main() {
	tempC := 18.0

	f := toFahrenheit(tempC)
	k := toKelvin(tempC)

	fmt.Printf("Celsius: %.1f\n", tempC)
	fmt.Printf("Fahrenheit: %.1f\n", f)
	fmt.Printf("Kelvin: %.2f\n", k)

	// reverse test
	fmt.Printf("Back to Celsius: %.1f\n", toCelsius(f))
}
