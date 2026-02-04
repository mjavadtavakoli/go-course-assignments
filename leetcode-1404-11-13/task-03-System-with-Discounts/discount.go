package order

func CalculateSubtotal(order Order) float64 {
	var subtotal float64

	for _, item := range order.Items {
		subtotal += float64(item.Quantity) * item.Price
	}

	return subtotal
}

func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
	subtotal := CalculateSubtotal(order)
	maxDiscount := 0.0

	for _, rule := range rules {
		if subtotal >= rule.MinAmount && rule.DiscountPercent > maxDiscount {
			maxDiscount = rule.DiscountPercent
		}
	}

	order.Discount = maxDiscount
	order.Total = subtotal - (subtotal * maxDiscount / 100)

	return order
}
