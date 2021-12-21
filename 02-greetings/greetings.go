package greetings

import (
	"time"
)

// Get greetings based on time. Return the greetings in Japanese
func greetings_by_time() string {
	now := time.Now()
	greeting := "こんにちは"
	currentHour := now.Hour()
	if currentHour > 4 && currentHour < 12 {
		greeting = "おはようございます"
	} else if currentHour > 18 || currentHour < 4 {
		greeting = "こんばんは"
	}
	return greeting
}

// Returns a greeting to a person in Japanese
func Hello(name string) string {
	message := greetings_by_time() + name + "様"
	return message
}
