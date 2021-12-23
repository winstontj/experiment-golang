package greetings

import (
	"errors"
	"math/rand"
	"time"
)

// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixMilli())
}

// Get a random value to get your luck
func get_luck() string {
	lucks := []string{
		"Good Luck",
		"Bad Luck",
		"Normal Luck",
		"Terrible Luck",
		"Amazing Luck",
	}
	return lucks[rand.Intn(len(lucks))]
}

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
func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("name is empty")
	}
	message := greetings_by_time() + name + "様. \nYou have " + get_luck()
	return message, nil
}

// A greeting that accept multiple names and return a map of name and greeting
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	if len(names) == 0 {
		return nil, errors.New("please provide name")
	}
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}
