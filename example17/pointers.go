package main

import "fmt"

// ival int is pass-by-copy
func zeroval(ival int) {
	ival = 0
}

// iptr *int is pass-by-reference
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// doesn't change i because ival is a copy of i
	zeroval(i)
	fmt.Println("zeroval:", i)

	// changes i because it is a pointer to i passed in
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// pointers can be printed
	fmt.Println("pointer:", &i)
}
