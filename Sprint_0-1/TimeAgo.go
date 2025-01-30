package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	d := time.Now().Sub(pastTime)
	years := int(d.Hours() / 24 / 365)
	months := int(d.Hours()/24/30) % 12
	days := int(d.Hours()/24) % 30
	hours := int(d.Hours()) % 24
	minutes := int(d.Hours()) % 60
	secondes := int(d.Seconds()) % 60

	if years >= 1 {
		return fmt.Sprintf("%d years ago", years)
	}
	if months >= 1 {
		return fmt.Sprintf("%d months ago", months)
	}
	if days >= 1 {
		return fmt.Sprintf("%d days ago", days)
	}
	if hours >= 1 {
		return fmt.Sprintf("%d hours ago", hours)
	}
	if minutes >= 1 {
		return fmt.Sprintf("%d minutes ago", minutes)
	}
	if secondes >= 1 {
		return fmt.Sprintf("%d secondes ago", secondes)
	}
	return ""
}

func main() {
	pastTime := time.Now().Add(-2 * time.Hour)
	fmt.Println(TimeAgo(pastTime))

}
