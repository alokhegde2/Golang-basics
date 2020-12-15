package main

import "fmt"

func main() {
	fmt.Println()
	//Arrays : Arrays have to be fixed length
	// var fruitArr [2]string

	//Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declaring and assigning
	// fruitArr := [2]string{"Apple", "orange"}

	// b := [5]int{1, 2, 3, 4, 5}

	// fmt.Println(b)
	// //outputing single array element
	// fmt.Println(fruitArr[1])

	//slices
	fruitSlice := []string{"Apple", "orange", "grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
	
}
