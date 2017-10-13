package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Scanner is like an input iterator
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Translate what is read to upoer case and output it
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
