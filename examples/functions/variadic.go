package main

import (
	"fmt"
)

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func dropFirstLast(grades ...int) int {
	total := 0
	grades = grades[1 : len(grades)-1]
	for _, grade := range grades {
		total += grade
	}
	return total / len(grades)
}

func main() {
	// Providing a Slice as an argument
	grades := []int{10, 8, 9, 10, 2}
	avg := dropFirstLast(grades...)
	fmt.Println("The average is:", avg)

}

func main1() {
	// Providing four arguments
	total := Sum(1, 2, 3, 4)
	fmt.Println("The Sum is:", total)

	// Providing three arguments
	total = Sum(5, 7, 8)
	fmt.Println("The Sum is:", total)

	// without arguments
	total = Sum()
	fmt.Println("The Sum is:", total)

	// Providing a Slice as an argument
	nums := []int{1, 2, 3, 4, 5}
	total = Sum(nums...)
	fmt.Println("The Sum is:", total)
}
