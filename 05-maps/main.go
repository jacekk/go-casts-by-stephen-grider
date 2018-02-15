package main

import (
	"fmt"
)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func main() {
	/*
		// way #1
		colors := make(map[string]string)
		colors["red"] = "#ff0000"
		colors["green"] = "#00ff00"
	*/

	/*
		// way #2
		var colors map[string]string
		colors = map[string]string{
			"red": "#ff0000",
		}
		colors["green"] = "#00ff00"
	*/

	// way #3
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"black": "#000000",
		"green": "#00ff00",
	}

	colors["white"] = "#ffffff"
	delete(colors, "blue")

	printMap(colors)
}
