package inventory

import "fmt"

func AddProduct(inventory map[int]Product, product Product) error {
	existingProduct, exists := inventory[product.ID]

	if exists {
		existingProduct.Quantity += product.Quantity
		inventory[product.ID] = existingProduct
	} else {
		inventory[product.ID] = product
	}

	return nil
}

func RemoveProduct(inventory map[int]Product, productID int) error {
	_, exists := inventory[productID]

	if !exists {
		return fmt.Errorf("product not found")
	}

	delete(inventory, productID)
	return nil
}

func CheckStock(inventory map[int]Product, productID int) (int, bool) {
	product, exists := inventory[productID]

	if !exists {
		return 0, false
	}

	return product.Quantity, true
}

func CalculateTotalValue(inventory map[int]Product) float64 {
	var total float64

	for _, product := range inventory {
		total += float64(product.Quantity) * product.Price
	}

	return total
}
