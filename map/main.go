package main

import "fmt"

func main() {
	bgColors := make(map[string]string)
	bgColors["white"] = "#ffffff"

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}

	delete(colors, "green")

	printMap(colors)
	printMap(bgColors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
