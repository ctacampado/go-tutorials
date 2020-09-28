package main

import (
	"fmt"
	"go-tutorials/internal/tempconv"
	"os"
	"strconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
