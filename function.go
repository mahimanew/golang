package main

import "fmt"

func main() {
	fmt.Println("Main function")
	greater()
	result := adder(3, 5)
	fmt.Println(result)
	fmt.Println(proAdder(1, 2, 3))

}

func greater() {
	fmt.Println("Greeter fucntion")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi this is your added values"
}
