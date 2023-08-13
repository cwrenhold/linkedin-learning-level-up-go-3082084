package main

import (
	"flag"
	"log"
)

// coin contains the name and value of a coin
type coin struct {
	name  string
	value float64
}

// coins is the list of values available for making change.
var coins = []coin{
	{name: "1 pound", value: 1},
	{name: "50 pence", value: 0.50},
	{name: "20 pence", value: 0.20},
	{name: "10 pence", value: 0.10},
	{name: "5 pence", value: 0.05},
	{name: "1 penny", value: 0.01},
}

// calculateChange returns the coins required to calculate the
func calculateChange(amount float64) map[coin]int {
	change := make(map[coin]int)

	for _, coin := range coins {
		// Note: Could also used an if here and then used division rather than nested for loops.
		for amount >= coin.value {
			amount -= coin.value
			change[coin]++
		}
	}

	return change
}

// printCoins prints all the coins in the slice to the terminal.
func printCoins(change map[coin]int) {
	if len(change) == 0 {
		log.Println("No change found.")
		return
	}
	log.Println("Change has been calculated.")
	for coin, count := range change {
		log.Printf("%d x %s \n", count, coin.name)
	}
}

func main() {
	amount := flag.Float64("amount", 0.0, "The amount you want to make change for")
	flag.Parse()
	change := calculateChange(*amount)
	printCoins(change)
}
