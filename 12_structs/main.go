package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting methode (value reciver)
func (p Person) greet() string {
	return "Hello My name is " + p.firstName + " " + p.lastName + " I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciver)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciver)
func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "f" {
		return
	} else {
		p.lastName = spouceLastName
	}
}

func main() {
	//Initializing person using structure
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	//Alternative
	person2 := Person{"Samantha", "Smith", "Boston", "f", 25}

	fmt.Println(person1)
	fmt.Println(person2)

	//to get single field
	fmt.Println(person1.firstName)

	//updating value
	// person1.age++

	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())
}
