package main

import "fmt"

func main() {

	// Declare Method 1

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	// Declare Method 2
	/*
		var colors map[string]string
	*/

	// Declare Method 3
	/*
		colors := make(map[string]string)
		colors["white"] = "#ffffff"
	*/

	// Delete key-value pairs
	/*
		delete(colors, "white")
	*/

	printMap(colors)
}

// Iterating over Maps
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
