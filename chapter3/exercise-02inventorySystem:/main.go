package main

import "fmt"

func main() {
	inventory := make(map[string]int)

	inventory["Laptop"] = 5
	inventory["Mouse"] = 10
	inventory["Keyboard"] = 7

	product := "Mouse"

	if qty, exists := inventory[product]; exists {
		if qty > 0 {
			inventory[product]--
		}
	}

	for name, qty := range inventory {
		fmt.Println(name, ":", qty)
	}
}
