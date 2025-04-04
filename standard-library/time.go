package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now", now)

	utc := time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
	fmt.Println("utc", utc)
	fmt.Println("utc local", utc.Local())

	fmt.Println("==========\n")

	formatDate := "2006-01-02 15:04:05"
	valueDate := "2025-01-10 01:01:01"

	timeFormat, err := time.Parse(formatDate, valueDate)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(timeFormat)
	}

	fmt.Println("==========\n time duration\n")

	duration1 := time.Second * 100
	duration2 := time.Millisecond * 100

	fmt.Println(duration1, duration2)

}
