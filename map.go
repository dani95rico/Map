// EXAMPLE OF USE: MAP IN GO LENGUAGE BY DANI95RICO

package main

import "fmt"

func main() {
	colors := map[string]string{
		"RED":   "#ff0000",
		"GREEN": "#4bf745",
	}
	colors["WHITE"] = "#fffff"

	delete(colors, "RED")

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for ", key, "is ", value)
	}
}
