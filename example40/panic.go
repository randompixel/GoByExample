package main

import "os"

func main() {
	panic("A problem, eh")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
