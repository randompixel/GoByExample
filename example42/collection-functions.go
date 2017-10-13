package main

import (
	"fmt"
	"strings"
)

/**
Find the first index position of a string in an array of strings
or -1 if not found
*/
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

/**
Is it in the array?
*/
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

/**
Do any strings in the array match the function?
*/
func Any(vs []string, f func(string) bool) interface{} {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/**
Do all strings in the array match the function?
*/
func All(vs []string, f func(string) bool) interface{} {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

/**
Filter array for items containing query in the passed in function
*/
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

/**
Returns new slice of array based on function passed in
*/
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	// Get position of "pear"
	fmt.Println(Index(strs, "pear"))

	// Does it include "grape"?
	fmt.Println(Include(strs, "grape"))

	// Does it have anything starting with "p"?
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// Does it have EVERYTHING starting with "p"?
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// Return anything with an "e" in it
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// Upper-case all items in array
	fmt.Println(Map(strs, strings.ToUpper))

}
