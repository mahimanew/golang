package main

import "fmt"

func main() {

	mahi := User{"Mahi", "mahimaabc@gmail.com", true, 10}
	fmt.Println(mahi)
	fmt.Println(mahi.name)
	fmt.Println(mahi.email)
	fmt.Println(mahi.age)
	fmt.Println(mahi.status)

}

type User struct {
	name   string
	email  string
	status bool
	age    int
}
