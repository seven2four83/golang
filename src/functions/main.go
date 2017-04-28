package main

import (
	"fmt"
	"time"
)

func main() {
	hourOfTheDay := time.Now().Hour()
	greeting := getGreeting(hourOfTheDay)
	greeting += ", Seven2Four"
	fmt.Println(greeting)
}

func getGreeting(hourOfTheDay int) string {
	if hourOfTheDay < 12 {
		return "Good Morning"
	} else {
		return "Good Evening"
	}
}
