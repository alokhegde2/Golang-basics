package main

import(
	"fmt"
)

func main(){
	//bool datatype : it is either true or false

	// var n bool = true
	// fmt.Printf("%v , %T",n,n)



	//numeric types : 
		//Integers
		// n :=42
		// fmt.Printf("%v , %T",n,n)


	//if we make the integer division it will give only integer number
		// a := 10
		// b := 3
		// fmt.Println(a/b)



		//floating number
		//We cannot able to do calculations b/w float32 and float64

		// var n float32 = 3.14   
		// n := 3.14	// this will assign it to the float64
		//n = 13.7e72
		// n = 2.1E14
		// fmt.Printf("%v , %T",n,n)

		//remainder operator is not availabe in float it only available in int


		//Complex number
		//there are two types of complex number
		//1:complex64 and complex128
		// var n complex64 = 1+2i
		// fmt.Printf("%v , %T\n",n,n)
		//operator available +,/,*,-


		//text type
		//1.string : any utf-8 letters
		//string will be read in the form of arrays
		// s := "this is the string"
		// fmt.Printf("%v , %T\n",s,s)
		//string concatination

		// s := "this is the string"
		// s2 := " this is also a string"
		// fmt.Printf("%v , %T\n",s+s2 , s+s2)

		//2.rune : it is utf-32 charector
		//rune value should be written in single quote
		var r rune = 'a'
		fmt.Printf("%v , %T\n",r,r)

}