package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	if start.Weekday() == 6 {
		return start.Add(+48 * time.Hour)
	}
	if start.Weekday() == 5 {
		return start.Add(+72 * time.Hour)
	} else {
		return start.Add(+24 * time.Hour)
	}
	return start
}

func main() {
	start := time.Date(2023, time.October, 6, 0, 0, 0, 0, time.UTC) // Friday
	fmt.Println(NextWorkday(start))
}
