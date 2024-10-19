package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Since(time.Now().Add(-2 * time.Hour))
	years := int(d.Hours() / 24 / 365)
	months := int(d.Hours()/24/30) % 12
	days := int(d.Hours()/24) % 30
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	if years > 1 {
		fmt.Printf("%d years ago", years)
	} else if years == 1 {
		fmt.Printf("%d year ago", years)
	} else if months > 1 {
		fmt.Printf("%d months ago", months)
	} else if months == 1 {
		fmt.Printf("%d month ago", months)
	} else if days > 1 {
		fmt.Printf("%d days ago", days)
	} else if days == 1 {
		fmt.Printf("%d day ago", days)
	} else if hours > 1 {
		fmt.Printf("%d hours ago", hours)
	} else if hours == 1 {
		fmt.Printf("%d hour ago", hours)
	} else if minutes == 1 || minutes > 1 {
		fmt.Printf("%d minutes ago", minutes)
	} else if seconds == 1 || seconds > 1 {
		fmt.Printf("%d seconds ago", seconds)
	}
	fmt.Printf("\n%02d:%02d:%02d:%02d:%02d:%02d\n", years, months, days, hours, minutes, seconds)
}
