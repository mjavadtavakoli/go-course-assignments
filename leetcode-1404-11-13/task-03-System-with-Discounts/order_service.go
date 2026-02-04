package order

func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
	for i, order := range orders {
		order.Status = "processing"

		order = ApplyDiscountRules(order, rules)

		if order.Total < 0 {
			order.Status = "cancelled"
		} else {
			order.Status = "completed"
		}

		orders[i] = order
	}

	return orders
}

func FilterOrdersByStatus(orders []Order, status string) []Order {
	var result []Order

	for _, order := range orders {
		if order.Status == status {
			result = append(result, order)
		}
	}

	return result
}

func CalculateOrderStatistics(orders []Order) map[string]interface{} {
	stats := make(map[string]interface{})

	totalRevenue := 0.0
	totalDiscount := 0.0
	completedOrders := 0

	for _, order := range orders {
		if order.Status == "completed" {
			completedOrders++
			totalRevenue += order.Total
			totalDiscount += order.Discount
		}
	}

	stats["total_orders"] = len(orders)
	stats["total_revenue"] = totalRevenue
	stats["completed_orders"] = completedOrders

	if completedOrders > 0 {
		stats["average_order_value"] = totalRevenue / float64(completedOrders)
	} else {
		stats["average_order_value"] = 0.0
	}

	stats["total_discount"] = totalDiscount

	return stats
}
