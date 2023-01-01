package main

import (
	tc "ex1/tempconv"
	"fmt"
)

func main() {
	var c tc.Celsius
	var f tc.Fahrenheit
	var k tc.Kelvin

	fmt.Println(c.String())
	fmt.Println(f.String())
	fmt.Println(k.String())
	fmt.Println(c == tc.Celsius(f))

	fmt.Println(tc.CToF(tc.BoilingC) == tc.BoilingF)
	fmt.Println(tc.FToC(tc.AbsoluteZeroF) == tc.AbsoluteZeroC)
	fmt.Println(tc.KToC(tc.AbsoluteZeroK) == tc.AbsoluteZeroC)
	fmt.Println(tc.CToK(tc.FreezingC) == tc.FreezingK)
	fmt.Println(tc.KToF(tc.FreezingK) == tc.FreezingF)
	fmt.Println(tc.FToK(tc.BoilingF) == tc.BoilingK)
}
