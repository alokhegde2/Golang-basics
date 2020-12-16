// Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Don't execute this")
	case true:
		fmt.Println("Execute this")

	}
}
