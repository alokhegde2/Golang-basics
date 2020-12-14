package main

import "fmt"

func main() {

	//Range is use to loop through arrays map and etc
	ids := []int{33, 76, 54, 23, 11, 2}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%v -ID: %v\n", i, id)
	}

	//not using index
	for _, id := range ids {
		fmt.Printf("ID: %v\n", id)
	}

	//Add ids togather
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%v : %v\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
