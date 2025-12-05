package main

import (
	"fmt"
	"time"
)

func main() {
	// Unix epoch
	// 00:00:00 UTC on January 1, 1970

	now := time.Now()
	unixTime := now.Unix()

	fmt.Println("current unix time: ", unixTime)

	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02"))



}
