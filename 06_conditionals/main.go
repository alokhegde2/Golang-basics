package main

import "fmt"

func main() {
	x := 15
	y := 10

	//You can add pracket to condition if(x<y){}
	//if else
	if x <= y {
		fmt.Printf("%v is less than or equal %v\n ", x, y)
	} else {
		fmt.Printf("%v is less than %v\n ", y, x)
	}

	//elseif
	color := "red"

	if color == "red" {
		fmt.Println("Colour is red")
	} else if color == "blue" {
		fmt.Println("Colour is blue")
	} else {
		fmt.Println("Colour is nt blue or red")
	}

	//Switch

	switch color {
	case "red":
		fmt.Println("Colour is red")
	case "blue":
		fmt.Println("Colour is blue")
	default:
		fmt.Println("Colour is not blue or red")
		
	}
}
