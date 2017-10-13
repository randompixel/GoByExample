package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Use crypto/rand for more secure randomness
func main() {

	// Create two random numbers between 0 > 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// Create a random float between 0.0 < 1.0
	fmt.Println(rand.Float64())

	// Create random float between 5.0 <= 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// The above are deterministic. Adding a seed makes it more
	// random
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// However, using a fixed seed still results in the same number
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	// .. no matter how often you see it
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
