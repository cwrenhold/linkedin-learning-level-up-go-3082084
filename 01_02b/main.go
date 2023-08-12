package main

import (
	"log"
	"time"
	"strings"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")

	for _, word := range words {
		var toPrint []string

		for index, char := range word {
			repeat := strings.Repeat(string(char), index + 1)
			toPrint = append(toPrint, repeat)
		}

		print(strings.Join(toPrint, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
