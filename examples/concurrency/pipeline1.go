package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 5; x++ {
			naturals <- x
			time.Sleep(time.Second)
		}
		fmt.Println("Going to close naturals channel")
		time.Sleep(time.Second)
		close(naturals)
		fmt.Println("Done closing natural channel")
		time.Sleep(time.Second)
	}()

	// squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
			time.Sleep(time.Second)
		}
		fmt.Println("Going to close squares channel")
		time.Sleep(time.Second)
		close(squares)
		fmt.Println("Done closing squares channel")
		time.Sleep(time.Second)
	}()

	for {
		fmt.Println(<-squares)
		time.Sleep(time.Second)
	}

}
