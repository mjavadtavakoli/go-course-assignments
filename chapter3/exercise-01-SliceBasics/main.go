package main

import "fmt"

func main() {
	cars := []string{"pride", "samand", "rana"}

	cars = append(cars, "l90")

	fmt.Println("Length:", len(cars))

	for i, item := range cars {
		fmt.Println(i, item)
	}
}
