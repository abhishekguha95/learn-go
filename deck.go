// This is the package declaration - every Go file must belong to a package
// 'main' is a special package that creates an executable program
package main

// Import statements bring in external packages we need
import (
	"fmt"     // For printing to console (formatted I/O)
	"os"      // For operating system interface (file operations)
	"strings" // For string manipulation functions
)

// Define a new type 'deck' which is a slice of strings
// This creates a custom type based on Go's built-in slice type
// Think of it as creating our own data type specifically for playing cards
type deck []string

// newDeck creates and returns a new deck of cards
// This is a function that returns our custom 'deck' type
func newDeck() deck {
	cards := deck{} // Initialize an empty deck using our custom type

	// Create slices (dynamic arrays) for card suits and values
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"} // All possible suits
	cardValues := []string{"Ace", "Two", "Three", "Four"}           // All possible values

	// Nested loops to create all possible card combinations
	// Outer loop: iterate through each suit
	for _, suite := range cardSuites {
		// Inner loop: for each suit, iterate through each value
		for _, value := range cardValues {
			// append() adds a new element to the slice
			// We're creating card names like "Ace of Spades"
			cards = append(cards, value+" of "+suite) // Add card to the deck
		}
	}

	return cards // Return the complete deck of 16 cards
}

// showCards prints each card in the deck with its index
// (d deck) is a receiver - this makes showCards a method on the deck type
// This means we can call it like: myDeck.showCards()
func (d deck) showCards() {
	// range gives us both index (i) and value (card) from the slice
	for i, card := range d {
		fmt.Println(i, card) // Print index and card value to console
	}
}

// deal splits the deck into two: a hand of 'size' cards and the remaining deck
// This function takes a deck and an integer, returns two decks
func deal(d deck, size int) (deck, deck) {
	// Slice syntax: d[:size] means "from start up to (but not including) size"
	hand := d[:size] // Take the first 'size' cards as the hand
	// d[size:] means "from index 'size' to the end"
	d = d[size:]   // The rest of the deck after dealing
	return hand, d // Return both the hand and the remaining deck
}

// convert deck to a string
// This method converts our deck slice into a single comma-separated string
func (d deck) toString() string {
	// strings.Join takes a slice of strings and joins them with a separator
	// []string(d) converts our custom 'deck' type back to a regular string slice
	return strings.Join([]string(d), ",")
}

// saveFile saves the deck to a file on disk
// This method takes a filename and returns an error (Go's way of error handling)
func (d deck) saveFile(filename string) error {
	// os.WriteFile writes data to a file
	// []byte(d.toString()) converts our string to bytes (what files store)
	// 0666 is the file permission (readable/writable by owner, group, and others)
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// readFile reads a deck from a file and returns it as a deck type
// This function takes a filepath string and returns a deck and an error
func readFile(filepath string) (deck, error) {
	// os.ReadFile returns two values: the file content as []byte and an error
	// In Go, we handle errors explicitly - this is a core Go principle
	cardFromDisk, err := os.ReadFile(filepath) // Use the filepath parameter, not hardcoded path

	// Check if there was an error reading the file
	// In Go, nil means "no error" - if err is not nil, there was an error
	if err != nil {
		// If there's an error, return an empty deck and the error
		// deck{} creates an empty deck, err passes the error up to the caller
		return deck{}, err
	}

	// Convert the byte slice to a string, then split by commas to recreate the deck
	// string(cardFromDisk) converts []byte to string
	// strings.Split separates the comma-separated string back into a slice
	deckFromFile := deck(strings.Split(string(cardFromDisk), ","))

	// Return the deck and nil (no error)
	return deckFromFile, nil
}
