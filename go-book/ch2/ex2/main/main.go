package main

import (
	uc "ex2/unitconv"
	"fmt"
)

func main() {
	var pa uc.Pascal = 101300
	var atm uc.Atmospheric = 1

	fmt.Println(pa.String())
	fmt.Println(atm.String())

	fmt.Println(uc.AtmToPa(atm))
	fmt.Println(uc.PaToAtm(pa))
}
