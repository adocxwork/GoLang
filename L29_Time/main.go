package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Type of currentTime : %T\n", currentTime)
	
	formatted := currentTime.Format("02-01-2006, Monday, 15:04:05, 3:04 PM")
	fmt.Println("Formatted time :", formatted)
	
	
	// layout same hona chahiye, date yaad rakhna hoga
	layout_string := "2006-01-02"
	dateStr := "2023-11-25"
	formattedTime, _ := time.Parse(layout_string, dateStr)
	fmt.Println(formattedTime)

	// add 1 more day in current time
	newDate := currentTime.Add(24 * time.Hour)
	fmt.Println("New date :", newDate)
	formattedNewDate := newDate.Format("2006/01/02 Monday")
	fmt.Println(formattedNewDate)
}