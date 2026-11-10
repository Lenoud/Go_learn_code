package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "pertice1: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}

	for line, fileCounts := range counts {
		n := 0
		filenames := make([]string, 0, len(fileCounts))
		for filename, fileCount := range fileCounts {
			n += fileCount
			filenames = append(filenames, filename)
		}
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, filenames)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][filename]++
	}
}
