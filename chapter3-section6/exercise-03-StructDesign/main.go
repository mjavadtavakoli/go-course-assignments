package main

import "fmt"

type OrderItem struct {
	ProductID int
	Name      string
	Price     float64
	Quantity  int
}

type Order struct {
	ID         int
	CustomerID string
	Items      []OrderItem
}

func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
}

func (o Order) CalculateTotal() float64 {
	var total float64
	for _, item := range o.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (o Order) PrintSummary() {
	fmt.Printf("Order #%d for Customer %s\n", o.ID, o.CustomerID)
	fmt.Println("Items:")
	for i, item := range o.Items {
		subtotal := item.Price * float64(item.Quantity)
		fmt.Printf("  %d. %s x%d @ $%.2f = $%.2f\n",
			i+1, item.Name, item.Quantity, item.Price, subtotal)
	}
	fmt.Printf("Total: $%.2f\n", o.CalculateTotal())
}

func main() {
	order := Order{
		ID:         1001,
		CustomerID: "cust-123",
	}

	order.AddItem(OrderItem{1001, "Laptop", 999.99, 1})
	order.AddItem(OrderItem{1002, "Mouse", 29.99, 2})
	order.AddItem(OrderItem{1003, "Keyboard", 79.99, 1})

	order.PrintSummary()
}
