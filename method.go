package main

import "fmt"

func main() {
	mahi := User{"mahi", "Mahi@gmail.com", "12345678", 20}
	mahi.getUser()
	mahi.getEmail()
	mahi.getAge()
	mahi.getMobile()
}

type User struct {
	uname  string
	email  string
	mobile string
	age    int
}

func (u User) getUser() {
	fmt.Println(u.uname)
}

func (u User) getEmail() {
	fmt.Println(u.email)
}

func (u User) getAge() {
	fmt.Println(u.age)
}

func (u User) getMobile() {
	fmt.Println(u.mobile)
}
