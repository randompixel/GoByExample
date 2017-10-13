package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// print struct
	fmt.Printf("%v\n", p)

	// print struct with field names
	fmt.Printf("%+v\n", p)

	// # adds a source-code representation of p
	fmt.Printf("%#v\n", p)

	// print the type of the value
	fmt.Printf("%T\n", p)

	// format boolean
	fmt.Printf("%t\n", true)

	// format standard base 10 decimal
	fmt.Printf("%d\n", 123)

	// format as binary number - 1110
	fmt.Printf("%b\n", 14)

	// print character of given integer - !
	fmt.Printf("%c\n", 33)

	// print hex of given integer - 1c8
	fmt.Printf("%x\n", 456)

	// format as basic float
	fmt.Printf("%f\n", 78.9)

	// format float in scientific notion = 1.234000e+08
	fmt.Printf("%e\n", 123400000.0)

	// format float in scientific notion - 1234e+08
	fmt.Println("%E\n", 123400000.0)

	// Basic string formatting - "string"
	fmt.Printf("%s\n", "\"string\"")

	// format string with Double-Quotes
	fmt.Printf("%q\n", "\"string\"")

	// print representation of pointer
	fmt.Printf("%p\n", &p)

	// add blank padding with number after % - |   12|  345|
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// add blank padding with number and precision after % - |  1.20|  3.45|
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// left justify the above
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Sprintf returns strings rather than printing them
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Write to stderr rather than stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
