package main

import "fmt"

func main() {
	products := []string{"Laptop", "Mouse", "Keyboard"}

	products = append(products, "Monitor")

	fmt.Printf("Total products: %d\n", len(products))

	for i, product := range products {
		fmt.Printf("%d. %s\n", i+1, product)
	}
}
