package main

import (
	"fmt"
	"os"
)

// Run this with go build, then execute the program with at
// least 3 arguments
func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
