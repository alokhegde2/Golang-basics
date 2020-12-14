// Using the following operators, write expressions and assign their values to variables:
// ==
// <=
// >=
// !=
// <
// >
// Now print each of the variables.

package main

import (
	"fmt"
)

func main() {
	a := (42 == 47)
	b := (42 <= 47)
	c := (42 >= 47)
	d := (42 != 47)
	e := (42 > 47)
	f := (42 < 47)
	fmt.Println(a, b, c, d, e, f)
}
