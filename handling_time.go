package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println(current_time)
	fmt.Println(current_time.Format("01-02-2006"))
	fmt.Println(current_time.Format("01-02-2006 Monday"))
	fmt.Println(current_time.Format("01-02-2006 15:04:05 Monday"))

	//for custom date
	current_date := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(current_date)
	fmt.Println(current_date.Format("01-02-2006 Monday"))
}
