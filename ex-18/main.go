// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	name := "James Bon"
	if name == "James Bond" {
		fmt.Println("Yooo, correct name!")
	} else if name == "John Wick" {
		fmt.Println("Correct,You are John Wick !!!!")
	} else {
		fmt.Println("You note the person we wants")
	}
}
