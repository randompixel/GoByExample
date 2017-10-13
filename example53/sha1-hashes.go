package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	s := "sha1 this string"
	// Get a new Hash stream
	h := sha1.New()

	// Put string bytes onto underlying data stream
	h.Write([]byte(s))

	// Get bytes as a slice by applying empty to end
	bs := h.Sum(nil)

	// Convert byte-stream to Hex
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
