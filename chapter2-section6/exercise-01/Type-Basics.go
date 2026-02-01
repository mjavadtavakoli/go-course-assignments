package main

import "fmt"

func main() {
	productName := "Smartphone"
	price := 599.50
	stock := 30
	available := stock > 0

	fmt.Println("Product:", productName)
	fmt.Println("Price:", price)
	fmt.Println("Stock:", stock)
	fmt.Println("Available:", available)
}



//solution:

/*package main

import "fmt"

func main() {
    // Product name (string)
    var productName string = "Laptop"

    // Price (float64)
    var price float64 = 999.99

    // Quantity in stock (int)
    var quantityInStock int = 15

    // Is available (bool)
    var isAvailable bool = true

    fmt.Printf("Product: %s\n", productName)
    fmt.Printf("Price: $%.2f\n", price)
    fmt.Printf("Stock: %d units\n", quantityInStock)
    fmt.Printf("Available: %t\n", isAvailable)
}*/
