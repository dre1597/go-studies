package main

import (
	"fmt"
	"math/rand"
	"time"
)

func throwCoin() string {
	rand.Seed(time.Now().UTC().UnixNano())

	results := []string{
		"heads",
		"tails",
		"heads",
		"tails",
		"heads",
		"tails",
		"heads",
		"tails",
		"middle",
	}

	return results[rand.Int()%len(results)]
}

func main() {
	switch throwCoin() {
	case "heads":
		fmt.Println("heads")
	case "tails":
		fmt.Println("tails")
	default:
		fmt.Println("middle")
	}
}
