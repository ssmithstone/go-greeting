package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = "[Luke Skywalker]"
	}
	fmt.Printf("Hello, %s!\n", name)
	printCurrentTime()
}

func printCurrentTime() {
	t := time.Now()
	fmt.Printf("Current time: %d\n", t.Unix()) // current time in seconds since 1970
	fmt.Printf("Current Formatted time: %s\n", t.Format(time.RFC822))
	hour, m, sec := t.Clock() // localtime clock
	fmt.Printf("Hour: %d\nMinute: %d\nSecond: %d\n", hour, m, sec)
	year, month, day := t.Date() // localtime date
	fmt.Printf("Year: %d\nMonth: %d\nDay: %d\n", year, month, day)
}
