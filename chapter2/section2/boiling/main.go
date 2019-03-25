// Boiling prints the boiling point of water.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	// the degree symbol is typed with <optioin>-<shift>-8
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
