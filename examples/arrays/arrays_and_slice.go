package main

import "fmt"

// reverse reverses a slice of ints in place

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func sumArray(s []int) int {
	total := 0
	for i := 0; i < len(s); i++ {
		total += s[i]
	}
	return total
}
func sumArrays(a, b []int) []int {
	var sums []int
	for i := range a {
		sums = append(sums, b[i]+a[i])
	}
	return sums
}
func main() {
	a := [...]int{10, 20, 30}
	b := [...]int{1, 2, 3}

	fmt.Println(sumArray(a[:]))
	fmt.Println(sumArrays(a[:], b[:]))
	reverse(a[:])
	fmt.Println(a)
}
