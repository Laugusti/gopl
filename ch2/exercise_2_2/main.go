package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Laugusti/gopl/ch2/exercise_2_2/lengthconv"
	"github.com/Laugusti/gopl/ch2/exercise_2_2/volumeconv"
	"github.com/Laugusti/gopl/ch2/exercise_2_2/weightconv"
	"github.com/Laugusti/gopl/ch2/tempconv"
)

func main() {
	if len(os.Args) <= 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			val, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse float: %v\n", err)
				os.Exit(1)
			}
			printConversion(val)
		}

	} else {
		for _, arg := range os.Args[1:] {
			val, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse float: %v\n", err)
				os.Exit(1)
			}
			printConversion(val)
		}
	}
}

func printConversion(val float64) {
	f := tempconv.Fahrenheit(val)
	c := tempconv.Celsius(val)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	lb := weightconv.Pound(val)
	kg := weightconv.Kilogram(val)
	fmt.Printf("%s = %s, %s = %s\n", lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))

	mi := lengthconv.Mile(val)
	km := lengthconv.Kilometer(val)
	fmt.Printf("%s = %s, %s = %s\n", mi, lengthconv.MiToKm(mi), km, lengthconv.KmToMi(km))

	gal := volumeconv.Gallon(val)
	l := volumeconv.Liter(val)
	fmt.Printf("%s = %s, %s = %s\n", gal, volumeconv.GalToL(gal), l, volumeconv.LToGal(l))
}
