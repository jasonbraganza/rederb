package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// The original Python file was just an imperative set of instructions
// Since I cannot replicate that here, I hope to just build a main function with calls to
// what I think are obvious disparate functions

func main() {
	// First things first. Load up any variables you need from environment files.
	// Put in a comma seperated list of env files to load
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment file(s)")
	}
	my_name := os.Getenv("NAME")
	artist1 := os.Getenv("ARTIST1")
	fmt.Printf("My name is %s and my favourite modern musician is %s\n", my_name, artist1)
}
