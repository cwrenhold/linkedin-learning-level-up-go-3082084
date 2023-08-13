package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// String returns a string representation of a raffleEntry
func (r raffleEntry) String() string {
	return r.Name
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var entries []raffleEntry

	// Load the file to a byte slice
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON byte slice to a slice of raffleEntry structs
	err = json.Unmarshal(data, &entries)

	if err != nil {
		log.Fatal(err)
	}

	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
