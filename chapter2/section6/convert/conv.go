package convert

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius((k - 273.15)) }

// KToF converts a Kelvin temperature to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// FToK converts a Fahrenheit temperature to Kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// FToM converts feet to meters
func FToM(f Feet) Meters { return Meters(f / 3.281) }

// MToF converts meters to feet
func MToF(m Meters) Feet { return Feet(m * 3.281) }

// PToK converts pounds to kilograms
func PToK(p Pounds) Kilograms { return Kilograms(p / 2.205) }

// KToP converts kilograms to pounds
func KToP(k Kilograms) Pounds { return Pounds(k * 2.205) }
