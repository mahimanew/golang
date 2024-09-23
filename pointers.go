package main

import "fmt"

func main() {
	fmt.Println("Pointers sample")

	var ptr *int
	fmt.Println(ptr)  // print nil value
	fmt.Println(&ptr) // print the address of ptr

	mynum := 23
	ptr = &mynum
	fmt.Println(ptr)  // print the address of mynum
	fmt.Println(*ptr) // print mynum value

	*ptr = *ptr + 2 // Added +2 with Mynum
	fmt.Println(*ptr)

}
