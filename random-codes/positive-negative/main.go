package main

import "fmt"

func positiveNegative(x int) string {
	switch {
	case x > 0:
		return "Positive"
	default:
		return "Zero"
	case x < 0:
		return "Negative"
	}
}

func main() {
	fmt.Println(positiveNegative(1))
	fmt.Println(positiveNegative(-1))
	fmt.Println(positiveNegative(0))
}
