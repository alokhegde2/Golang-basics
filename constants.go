package main

import(
	"fmt"
)

// const a int16 = 27
const a = iota  //which is for enumarated value

func main(){
	//naming 

	// const myConst int = 42
	// fmt.Printf("%v , %T\n",myConst,myConst)

	// const myConst float32 = 42.89
	// fmt.Printf("%v , %T\n",myConst,myConst)

	// const a int = 14
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Printf("%v \n",a)
	// fmt.Printf("%v \n",b)
	// fmt.Printf("%v \n",c)
	// fmt.Printf("%v \n",d)

	fmt.Printf("%v , %T\n",a,a)
}