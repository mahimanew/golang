package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices example")
	var fruitList = []string{"Apple", "Banana", "Orange"}
	fmt.Println(fruitList)

	fmt.Printf("Type of fruitList is %T ", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:])
	fmt.Println(fruitList[1:3])

	var number [4]int
	number[0] = 10
	number[1] = 20
	number[2] = 30
	fmt.Println(number)

	num := make([]int, 4)
	num[0] = 100
	num[1] = 200
	num[2] = 300
	fmt.Println(num)

	num[3] = 400
	fmt.Println(num)

	num = append(num, 10, 20, 20)
	fmt.Println(num)

	sort.Ints(num)
	fmt.Println(num)

}
