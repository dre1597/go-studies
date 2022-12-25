package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// not 100% correct solution for the exercise, but part of

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func main() {
	args := os.Args

	startEcho1 := time.Now()
	echo1(args)
	durationEcho1 := time.Since(startEcho1)
	fmt.Printf("func echo1: %s\n", durationEcho1)

	startEcho2 := time.Now()
	echo2(args)
	durationEcho2 := time.Since(startEcho2)
	fmt.Printf("func echo2: %s\n", durationEcho2)

	startEcho3 := time.Now()
	echo3(args)
	durationEcho3 := time.Since(startEcho3)
	fmt.Printf("func echo3: %s\n", durationEcho3)
}
