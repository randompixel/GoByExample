package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // this stops the loop below

	for elem := range queue {
		fmt.Println(elem)
	}
}
