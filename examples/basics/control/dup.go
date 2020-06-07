package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup3() map[string]int {
	// dup3 finds duplicates with all the data held in memory at once
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	return counts
}

type lineInfo struct {
	count int
	files []string
}

func dup4() map[string]lineInfo {
	// in addition to printing the count of duplicates
	// also prints the files in which the dups are found
	counts := make(map[string]lineInfo)
	files := os.Args[1:]
	// doesn't take only stdin, needs to have list of files as input
	for _, arg := range files {
		f, _ := os.Open(arg)
		countLines2(f, counts)
	}
	for line, entry := range counts {
		fmt.Printf("%d\t%s\n", entry, line)

	}
	return counts
}

func countLines2(f *os.File, counts map[string]lineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		t := counts[input.Text()]
		t.count += 1
		t.files = append(t.files, f.Name())
		counts[input.Text()] = t
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup2() map[string]int {
	counts := make(map[string]int)
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
		}
	}
	return counts
}

func dup1() map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Printf("Counts is %v\n", counts)
	return counts
}
func main() {
	counts := dup3()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
