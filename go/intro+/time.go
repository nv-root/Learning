package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current time:", time.Now())

	specificTime := time.Date(2025, time.December, 5, 12, 0, 0, 0, time.UTC)
	// fmt.Printf("%T\n", specificTime)
	fmt.Println("Specific time: ", specificTime)


	// parsing time based on reference Mon Jan 2 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2025-12-05")
	// fmt.Printf("%T\n", parsedTime)
	fmt.Println("ParsedTime: ", parsedTime)

	parsedTime2, _ := time.Parse("06-01-02", "26-12-05")
	fmt.Println(parsedTime2)

	fmt.Println("Formatted time:", time.Now().Format("Monday 2006-01-02 15:04:05"))


	// adding 24 hours
	oneDayLater := time.Now().Add(time.Hour * 24)
	fmt.Println("24 hours later: ", oneDayLater)

	loc, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Println(time.Now().In(loc))

	// comparing time
	fmt.Println(time.Now().Before(oneDayLater))


	layout := "Jan 02, 2006 03:04 PM"
	t, _ := time.Parse(layout, "Dec 05, 2025 01:10 PM")
	fmt.Println(t)
}
