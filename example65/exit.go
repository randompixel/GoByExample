package main

import (
	"fmt"
	"os"
)

func main() {

	// defer doesn't get run when exiting, so ! will never show
	defer fmt.Println("!")


	// Using `go run`, exit code will be 1 as it is `run`'s exit code
	// to exit with 3, the program has to be built with `go build`
	os.Exit(3)
}
