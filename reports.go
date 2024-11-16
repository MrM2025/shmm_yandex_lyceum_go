package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	ReportID int
	Date     string
	User
}

func CreateReport(user User, reportDate string) Report {
	var r Report
	r.Date = reportDate
	r.ReportID = time.Now().Second()
	r.User = user
	return r
}

func PrintReport(report Report) {
	fmt.Println(report.Date)
	fmt.Println(report.User)
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var r []Report
	for _, user := range users {
		r = append(r, CreateReport(user, time.Now().Format("2006-01-02")))
	}
	return r
}

func main() {
	var users []User

	for i := 1; i <= 3; i++ {
		users = append(users, User{ID: i, Name: "Иван", Email: "ivan@example.com", Age: 30 + i})
	}

	reportDate := time.Now().Format("2006-01-02")
	reports := GenerateUserReports(users, reportDate)

	for _, report := range reports {
		PrintReport(report)
	}
}
