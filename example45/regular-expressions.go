package main

import (
	"fmt"
	"regexp"
	"bytes"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	// Match by Regular Expression
	fmt.Println(r.MatchString("peach"))

	// Find matches for Regular Expression
	fmt.Println(r.FindString("peach punch"))

	// Return array of start and end indexes of match
	fmt.Println(r.FindStringIndex("peach punch"))

	// Returns array of whole pattern and sub-pattern matches ()
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Returns array of indexes where subpatterns match [0 5 1 3]
	fmt.Println(r.FindStringSubmatchIndex("peach"))

	// Return array of all matches, -1 = non-limiter
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// As above but only return first two
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Work with bytes rather than strings
	fmt.Println(r.Match([]byte("peach")))

	// MustCompile works with constants because it only has a single return
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Replace strings
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Replace with a custom function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
