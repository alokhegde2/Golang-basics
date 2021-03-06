// Building on the code from the previous example
// at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// eg:

// in func main
// this should already be done
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”
// now do this
// now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// then use the “=” operator to ASSIGN that value to “y”
// print out the value stored in “y”
// print out the type of “y”

package main

import (
	"fmt"
)

type hotdog int //here we are creating our own datatype hotdog which is of integer

var x hotdog //here we are declaring variable x which of datatype hotdog

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x) //here we can't do y=x to assign the values of x to y because both are in diff datatype x in hotdog and y in int
	fmt.Println(y)
	
	fmt.Printf("%T\n", y)
}
