package main

import (
	"strconv"
	"fmt"
)

func main() {

	// Parse float with 64 bits of precision
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Parse int as base10 inside 64bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// Works with hex -> dec conversions
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Parse unsigned int as base10 inside 64bits
	// setting to negative gives zero
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Shortcut function for base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Throw error on bad input (ie: it's not an int string)
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
