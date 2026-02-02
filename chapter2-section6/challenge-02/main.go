package main

import (
    "fmt"
)

type Money int64

// Constructor
func NewMoney(dollars float64) Money {
    return Money(dollars * 100)
}

func (m Money) Dollars() float64 {
    return float64(m) / 100
}

func (m Money) String() string {
    if m < 0 {
        return fmt.Sprintf("-$%.2f", float64(-m)/100)
    }
    return fmt.Sprintf("$%.2f", float64(m)/100)
}

func (m Money) Add(other Money) Money {
    return m + other
}

func (m Money) Subtract(other Money) Money {
    return m - other
}

func (m Money) Multiply(factor int) Money {
    return m * Money(factor)
}

func main() {
    price := NewMoney(29.99)
    quantity := 3
    
    total := price.Multiply(quantity)
    discount := NewMoney(10.00)
    final := total.Subtract(discount)
    
    fmt.Printf("Unit price: %s\n", price)
    fmt.Printf("Quantity: %d\n", quantity)
    fmt.Printf("Subtotal: %s\n", total)
    fmt.Printf("Discount: %s\n", discount)
    fmt.Printf("Final: %s\n", final)
}