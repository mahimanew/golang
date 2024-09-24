package main

import "fmt"

func main() {
	fmt.Println("Arrays programming")
	var furitList [4]string
	furitList[0] = "Apple"
	furitList[1] = "Banana"
	furitList[2] = "Orange"

	fmt.Println(furitList)
	fmt.Println(len(furitList))

	var vegList = [5]string{"onion", "potato", "brinjal"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))

}
