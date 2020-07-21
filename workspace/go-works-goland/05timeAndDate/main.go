package main

import (
	"fmt"
	"time"
)

func main() {
	t:= time.Date(2020,time.July, 20,10,10,10,10,time.UTC)
	fmt.Printf("Time is %s \n", t)

	now:= time.Now();
	fmt.Printf("Time Now is : %s\n", now)

	// month
	fmt.Println("Month : ", t.Month())

	// Day
	fmt.Println("Day: " , t.Day())

	// Week Day
	fmt.Println("Week Day :", t.Weekday())

	tomorrow:=t.AddDate(0,0,1)
	fmt.Printf("Tomorrow: %s\n", tomorrow)

	fmt.Printf("Tomorrow Date: %v, Month %v\n", tomorrow.Day(), tomorrow.Month())

	fmt.Println("RFC850 : ", tomorrow.Format(time.RFC850))
	fmt.Println("RFC1123 : ",tomorrow.Format(time.RFC1123))

	// 7/22/2020
	fmt.Println(tomorrow.Format("Jan 2, 2006 at 3:04pm (MST)"))

	fmt.Println(tomorrow.Format("Monday, Jan"))

}