package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type ProductStore struct {
	products map[int]Product
	nextID   int
}

func NewProductStore() *ProductStore {
	return &ProductStore{
		products: make(map[int]Product),
		nextID:   1001,
	}
}

func (ps *ProductStore) Add(name string, price float64, stock int) int {
	id := ps.nextID
	ps.products[id] = Product{id, name, price, stock}
	ps.nextID++
	return id
}

func (ps *ProductStore) SearchByName(query string) []Product {
	query = strings.ToLower(query)
	results := []Product{}

	for _, product := range ps.products {
		if strings.Contains(strings.ToLower(product.Name), query) {
			results = append(results, product)
		}
	}

	return results
}

func (ps *ProductStore) FilterByPriceRange(min, max float64) []Product {
	results := []Product{}

	for _, product := range ps.products {
		if product.Price >= min && product.Price <= max {
			results = append(results, product)
		}
	}

	return results
}

func SortByPrice(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}

func main() {
	store := NewProductStore()

	store.Add("Gaming Laptop", 1299.99, 5)
	store.Add("Office Laptop", 799.99, 10)
	store.Add("Wireless Mouse", 29.99, 50)
	store.Add("Mechanical Keyboard", 89.99, 20)
	store.Add("USB Mouse", 19.99, 100)

	// Search
	results := store.SearchByName("laptop")
	fmt.Println("Search 'laptop':")
	for _, p := range results {
		fmt.Printf("  %s - $%.2f\n", p.Name, p.Price)
	}

	// Filter
	results = store.FilterByPriceRange(20, 100)
	SortByPrice(results)
	fmt.Println("\nPrice range $20-$100:")
	for _, p := range results {
		fmt.Printf("  %s - $%.2f\n", p.Name, p.Price)
	}
}
