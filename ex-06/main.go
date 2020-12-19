// Write a program that prints a number in decimal, binary, and hex

package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("%d\t\t%b\t\t%#x", x, x, x) //where %d for integer , %b for binary and %#x for hexadecimal
	
}
