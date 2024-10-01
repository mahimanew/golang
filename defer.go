package main

import "fmt"

func main() {
	defer fmt.Println("end")
	fmt.Println("Defer example")
	mydefer()

}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
