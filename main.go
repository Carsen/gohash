package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Declare constant 'goalHash' as string "password"
const goalString string = "password"

// Declare variable 'userHash' and 'goalHash' as byte slices
var userHash []byte
var goalHash []byte

// main function
func main() {
	// Declare variable 'userString' as string
	var userString string
	//
	// Opening Message
	fmt.Println("Welcome!")
	fmt.Print("Please enter a string: ")
	// User input into 'userString' variable
	fmt.Scanln(&userString)
	// Call hashAppend function
	hashAppend(userString, goalString)

	hex.EncodeToString(userHash)
	hex.EncodeToString(goalHash)

	if !bytes.Equal(userHash, goalHash) {
		fmt.Print("Incorrect!")
	} else if bytes.Equal(userHash, goalHash) {
		fmt.Print("Correct!")
	} else {
		fmt.Print("Hmm. Not sure how you did that!")
	}
	/*
	   // Print hashes in hexadecimal format
	   fmt.Printf("User Hash: %s\n", hex.EncodeToString(userHash))
	   fmt.Printf("Goal Hash: %s\n", hex.EncodeToString(goalHash))
	*/
}

// Function to hash strings and append them to bytes
func hashAppend(u string, g string) {
	// Shorthand 'h' to SHA256
	h := sha256.New()
	// Hash 'u' string to byte slice
	h.Write([]byte(u))
	// Append the sum of byte 'u' to userHash
	userHash = h.Sum(nil)
	// Reset the Hash
	h.Reset()
	// Hash 'g'
	h.Write([]byte(g))
	// Append the checksum of byte 'g' to goalHash
	goalHash = h.Sum(nil)
}
