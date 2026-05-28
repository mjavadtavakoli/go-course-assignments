package main

import (
	"context"
	"fmt"
	"time"
)

type order struct {
	price int
	id    string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	orderCh := receiveOrdersWithTimeout(ctx, "1", "2", "3", "4", "5")

	validOrderCh := validateOrderss(orderCh)
	for order := range validOrderCh {
		fmt.Println(order)
	}
}

func receiveOrdersWithTimeout(ctx context.Context, ids ...string) chan order {
	orderCh := make(chan order, len(ids))
	price := -2

	go func() {
		defer close(orderCh)
		for _, id := range ids {
			select {
			case <-ctx.Done():
				fmt.Println("Receive orders timed out or cancelled.")
				return
			case orderCh <- order{price: price, id: id}:
				price += 1
			}
		}
	}()
	return orderCh
}

func validateOrderss(in <-chan order) chan order {
	validOrderCh := make(chan order, 10)
	go func() {
		defer close(validOrderCh)
		for order := range in {
			if order.price >= 0 {
				validOrderCh <- order
			}
		}
	}()
	return validOrderCh
}
