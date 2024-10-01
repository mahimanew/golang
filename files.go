package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFile()
	readFile()

}

func writeFile() {
	content := "Hai hello"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	} else {
		length, err := io.WriteString(file, content)
		if err != nil {
			panic(err)
		}
		fmt.Println(length)
		file.Close()
	}
}

func readFile() {
	dataByte, err := ioutil.ReadFile("./myfile.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))
}
