package main

import (
	"fmt"
	"strings"
)

func swapIt(x, y string) (string, string) {
	return y, x
}

func main() {
	s := strings.Split(string("Ankit Gaur"), string(' '))
	x, y := s[0], s[1]
	fmt.Println("Before Swap:", x, y)

	x, y = swapIt(x, y)
	fmt.Println("After Swap:", x, y)
}
