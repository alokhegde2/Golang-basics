package main

import "fmt"

var name string = "Alok"

func main() {
	//MAIN TYPES
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	
	//creating variable by var variable
	// var name string = "Alok"
	var age int = 19
	var isCool bool = true
	isCool = false
	var size float32 = 2.3

	//shorthand operator
	mname := "Hegde"
	email := "alok@email.com"
	// size := 1.3

	lname, gmail := "Alok", "alok@example.com"

	fmt.Println(name, mname, age, isCool, size, email)
	fmt.Println(lname, gmail)

	//to get the type
	fmt.Printf("%T,%T,%T", name, age, isCool)

}
