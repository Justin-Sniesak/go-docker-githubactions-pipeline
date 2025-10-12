package main

import (
	"fmt"
	"time"
)

func greetingTimeDate() {
	name := "Justin"
	currentTime := time.Now()
	dateFormat := currentTime.Format("01/02/2006")
	timeFormat := currentTime.Format("03:04 PM")
	fmt.Println(name + ", the date is " + dateFormat + ", and the time is " + timeFormat + ".")
}

func main() {
	greetingTimeDate()
}
