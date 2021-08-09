// Exercise 2.2
package main

import (
	"fmt"
	"os"
	"strconv"

	"ch2/lengthconv"
	"ch2/tempconv"
	"ch2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Farenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

		m := lengthconv.Meter(t)
		ft := lengthconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToFt(m), ft, lengthconv.FtToM(ft))

		kg := weightconv.Kilogram(t)
		lbs := weightconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n", kg, weightconv.KgToLbs(kg), lbs, weightconv.LbsToKg(lbs))
	}
}
