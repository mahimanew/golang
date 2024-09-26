package main

import "fmt"

func main() {
	logincount := 25
	var result string

	if logincount < 10 {
		result = "Regular user"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	// one more syantx
	if num := 3; num >= 3 {
		fmt.Println("Number is greater than 3")
	} else {
		fmt.Println("Number is less than 3")
	}
}
