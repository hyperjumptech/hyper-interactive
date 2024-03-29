package main

import (
	"fmt"
	hyper_interactive "github.com/hyperjumptech/hyper-interactive"
	"time"
)

func main() {
	name := hyper_interactive.Ask("Whats your name ?", "Bruce Wayne", true)
	askTime := hyper_interactive.AskTime("Whats your time ?", time.Now(), false)
	options := []string{
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
	}
	choosen := hyper_interactive.Select("Please choose number 5", options, 1, 1, true)
	if choosen != 5 {
		fmt.Println("You should choose 5")
	} else {
		fmt.Printf("Thank you, %s, %s", name, askTime)
	}
}
