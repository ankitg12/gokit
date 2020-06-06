package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		fullname = "Ankit Gaur"
	)
	name := strings.Split(fullname, string(' '))
	firstname := name[0]
	// lastname := name[-1]

	for i := -5; i < 5; i++ {
		fmt.Println(i, "Hello, ", firstname)
	}
}
