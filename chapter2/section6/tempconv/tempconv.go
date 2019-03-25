// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius type for manipulating Celsius temperature values
type Celsius float64

// Fahrenheit type for manipulating Fahrenheit temperature values
type Fahrenheit float64

//Kelvin type for manipulating Kelvin temperature values
type Kelvin float64

const (
	// AbsoluteZeroC holds the value of absolute zero in Celsius
	AbsoluteZeroC Celsius = -273.15

	// FreezingC holds the value of freezing in Celsius
	FreezingC Celsius = 0

	// BoilingC holds the value of boiling in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
