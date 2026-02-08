package main

import "fmt"

func main() {
	products := []string{"pride", "samand", "rana"}

	products = append(products, "l90")

	fmt.Println("Length:", len(products))

	for i, item := range products {
		fmt.Println(i, item)
	}
}
