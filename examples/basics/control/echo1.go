package main

import (
	"fmt"
	"os"
)

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}
func echo1() string {
	s := os.Args[1]
	sep := " "
	for i := 2; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	return s

}
func main() {
	fmt.Println(echo2())
}
