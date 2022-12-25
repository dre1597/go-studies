package main

import (
	"bufio"
	"fmt"
	"os"
)

type Occurrence struct {
	files map[string]bool
	count int
}

func main() {
	counts := make(map[string]*Occurrence)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, oc := range counts {
		if oc.count > 1 {
			fmt.Printf("%d\t%s\n", oc.count, line)
			for f := range oc.files {
				fmt.Printf("\t%s\n\n", f)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]*Occurrence) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if _, ok := counts[text]; !ok {
			var oc Occurrence
			oc.files = make(map[string]bool)
			oc.count = 0
			counts[text] = &oc
		}
		counts[text].files[f.Name()] = true
		counts[text].count++
	}
}
