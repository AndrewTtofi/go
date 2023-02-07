package main

import "fmt"

func main() {
	/*alternative methods to declare a map
	var colors map[string]string
	colors := make(map[string]string)
	*/
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
	}
	//adding an additional value in the map
	colors["white"] = "#FFFFFF"
	//delete a value from a map
	delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for ", color, "is", hex)
	}
}
