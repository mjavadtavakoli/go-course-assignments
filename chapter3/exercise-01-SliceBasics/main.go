package main

import "fmt"

func main() {
	products := []string{"Laptop", "Mouse", "Keyboard"}

	products = append(products, "Monitor")

	fmt.Println("Length:", len(products))

	for i, item := range products {
		fmt.Println(i, item)
	}
}
