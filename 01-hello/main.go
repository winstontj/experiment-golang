package main

import (
	"fmt"

	quote "rsc.io/quote/v4"
	greetings "winstontj.github.com/greetings"
)

func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Amuro")
	fmt.Println(message)
}
