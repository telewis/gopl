// Package convert performs various conversions
package convert

import "fmt"

// Celsius type for manipulating Celsius temperature values
type Celsius float64

// Fahrenheit type for manipulating Fahrenheit temperature values
type Fahrenheit float64

//Kelvin type for manipulating Kelvin temperature values
type Kelvin float64

// Feet type for manipulating Feet values
type Feet float64

// Meters type for manipulating Meters values
type Meters float64

// Pounds type for manipulating Pounds values
type Pounds float64

// Kilograms type for manipulating Kilograms values
type Kilograms float64

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
func (f Feet) String() string       { return fmt.Sprintf("%g ft", f) }
func (m Meters) String() string     { return fmt.Sprintf("%g m", m) }
func (p Pounds) String() string     { return fmt.Sprintf("%g lb", p) }
func (k Kilograms) String() string  { return fmt.Sprintf("%g kg", k) }
