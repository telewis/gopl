// cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"../convert"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := convert.Fahrenheit(t)
		c := convert.Celsius(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, convert.FToC(f), c, convert.CToF(c), c, convert.CToK(c))

		ft := convert.Feet(t)
		m := convert.Meters(t)

		fmt.Printf("%s = %s, %s = %s\n", ft, convert.FToM(ft), m, convert.MToF(m))

		p := convert.Pounds(t)
		k := convert.Kilograms(t)

		fmt.Printf("%s = %s, %s = %s\n", p, convert.PToK(p), k, convert.KToP(k))
	}
}
