package main

import "fmt"

func main() {
	days := []string{"sun", "mon", "tue", "wed", "thr", "fri", "sat"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	for _, day := range days {
		fmt.Println(day)
	}

	//while
	num := 1

	for num <= 10 {
		fmt.Println(num)
		num++
	}

	num = 1
	for num <= 10 {

		if num == 5 {
			num++
			continue
		}
		fmt.Println(num)
		num++

	}

	num = 1
	for num <= 10 {

		if num == 5 {
			break
		}
		fmt.Println(num)
		num++

	}

	num = 1
	for num <= 10 {

		if num == 4 {
			goto mylabel
		}
		fmt.Println(num)
		num++
	}

mylabel:
	fmt.Println("This is example for go to concept")

}
