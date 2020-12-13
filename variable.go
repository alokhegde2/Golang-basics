package main

import (
	"fmt"
)

//global declaration

// var actorname string = "Alok Hegde"
// var companion string = "Sara"
// var doctorNumber int = 3

//above declarations are also written as , by this we can use the var only once

var(
	actorname string = "Alok Hegde"
	companion string = "Sara"
	doctorNumber int = 3
)


func main(){

//Declaring inside the main function

	// var i int
	// i=42
	// var j float32 = 27
	// k:=99
	// var x string ="Alok"

	//This print statement is similar to the c programming 

	// fmt.Printf("%v",actorname)

	//we can also use Println to print the single line

	// fmt.Println("The number of doctor is : ",doctorNumber)


	//Type conversion : Converting the variable from one type to another

	var i int = 42
	fmt.Printf("%v ,%T\n",i,i)	//where %T is used to check the type of the variable

	var j float32
	j = float32(i)	//here we are converting variable i from integer to float32
	fmt.Printf("%v ,%T\n",j,j) 
}