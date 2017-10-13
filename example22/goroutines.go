package main

import "fmt"

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Blocking function call
	f("direct")

	// Background Routine
	go f("goroutine")

	// Background Routine - Annon Function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input) // Force an input before exit
	fmt.Println("Done")
}
