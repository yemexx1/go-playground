package main

import (
	"fmt"
	"time"
	"errors"
	"os"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		err := errors.New("Too early or greetings!")
		return message, err
	}

	if hour < 12 {
		message = "Good Morning"
	} else if hour < 18 {
		message = "Good Afternoon"
	} else {
		message = "Good Evening"
	}

	return message, nil
}
