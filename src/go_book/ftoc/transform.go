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
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

	// Celsius to Fahrenheit
	fahrenheit := CtoF(-40)
	fmt.Printf("%g°F\n", fahrenheit)

	// fahrenheit to celsius
	celsius := fToC(32)
	fmt.Printf("%g°C\n", celsius)

}

func fToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
