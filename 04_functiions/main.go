package main

import (
	"fmt"
)

//infunction declaration first string is the datatype of name and second one is the return type If you have a return type only give the return types
// datatype or else keep it blank

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Alok"))
	fmt.Println(getSum(3, 4))
	
}
