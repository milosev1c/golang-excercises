package main

import (
	"course/ch2/tempconv/tempconv"
	"fmt"
)

func main() {
	fmt.Printf("%g", tempconv.Celsius(0))
	fmt.Printf("C to F: %g\n", tempconv.CToF(0))
	fmt.Printf("F to C: %g\n", tempconv.FToC(0))
	fmt.Printf("C to K: %g\n", tempconv.CToK(0))
	fmt.Printf("K to C: %g\n", tempconv.KToC(0))
	fmt.Printf("F to K: %g\n", tempconv.FToK(0))
	fmt.Printf("K to F: %g\n", tempconv.KToF(0))
}
