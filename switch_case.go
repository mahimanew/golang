package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the value")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 0)
	switch num {
	case 1:
		fmt.Println("You entered one")
		fallthrough
	case 2:
		fmt.Println("You entered two")
	default:
		fmt.Println("Entered number is not in the list")
	}
}
