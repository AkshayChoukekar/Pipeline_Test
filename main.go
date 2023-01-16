package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	year, month, date := time.Now().Date()

	dateString := strconv.Itoa(date) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(year)

	fmt.Println("Hello, have a good day! Todays Date is %v", dateString)
}
