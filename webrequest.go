package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://automationexercise.com/api/productsList"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	response.Body.Close()

}
