package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	// Read entire file into memory and check for errors
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Individual file operaion component parts
	f, err := os.Open("/tmp/dat")
	check(err)

	// Create a byte holder of up to 5 chars and read them in
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// Jump to position in file and read in 2 chars
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// Helper functions for certain ops
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Seek operates as rewind
	_, err = f.Seek(0, 0)
	check(err)

	// Bufio is enhanced buffered file reader
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close file when done, usually after opening with `defer`
	f.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
