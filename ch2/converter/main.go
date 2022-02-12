package main

import (
	"course/ch2/converter/conv"
	"course/ch2/tempconv/tempconv"
	"flag"
	"fmt"
)

func main() {
	var from = flag.String("from", "ft", "Что переводить")
	var to = flag.String("to", "m", "Куда переводить")
	var val = flag.Float64("val", 1.0, "Значение")
	flag.Parse()
	fmt.Println(*from, *to)
	switch {
	case *from == "m" && *to == "ft":
		fmt.Println(conv.MtoFt(conv.Meter(*val)))
	case *from == "ft" && *to == "m":
		fmt.Println(conv.FtToM(conv.Foot(*val)))
	case *from == "kg" && *to == "p":
		fmt.Println(conv.KgToP(conv.Kilogram(*val)))
	case *from == "p" && *to == "kg":
		fmt.Println(conv.PToKg(conv.Pound(*val)))
	case *from == "dK" && *to == "dF":
		fmt.Println(tempconv.KToF(tempconv.Kelvin(*val)))
	case *from == "dF" && *to == "dK":
		fmt.Println(tempconv.FToK(tempconv.Fahrenheit(*val)))
	case *from == "dC" && *to == "dK":
		fmt.Println(tempconv.CToK(tempconv.Celsius(*val)))
	case *from == "dK" && *to == "dC":
		fmt.Println(tempconv.KToC(tempconv.Kelvin(*val)))
	case *from == "dF" && *to == "dC":
		fmt.Println(tempconv.FToC(tempconv.Fahrenheit(*val)))
	case *from == "dC" && *to == "dF":
		fmt.Println(tempconv.CToF(tempconv.Celsius(*val)))
	}
	return
}
