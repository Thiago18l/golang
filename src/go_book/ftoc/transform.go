// conversion fahrenheit to celsius.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsolutZeroC Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, FToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, FToC(boilingF))

	// Celsius to Fahrenheit
	fahrenheit := CtoF(-40)
	fmt.Printf("%g°F\n", fahrenheit)

	// fahrenheit to celsius
	celsius := FToC(32)
	fmt.Printf("%g°C\n", celsius)

}

// fToC convert fahrenheit to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CtoF convert celsius to fahrenheit
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
