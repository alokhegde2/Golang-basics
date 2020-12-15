package main

import "fmt"

func main() {
	a := 5
	b := &a //assigning b to pointer to a , b value will be the memory address

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read values from address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change val with pointer
	*b = 10
	fmt.Println(a)
	
}
