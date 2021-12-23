package main

import (
	"fmt"
	"log"

	quote "rsc.io/quote/v4"
	greetings "winstontj.github.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())
	names := []string{
		"Eren",
		"Levi",
		"Rudy",
		"Emilia",
		"Jolyne",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for name, greeting := range messages {
		fmt.Println("Name: " + name)
		fmt.Println(greeting)
		fmt.Println("==============")
	}
}
