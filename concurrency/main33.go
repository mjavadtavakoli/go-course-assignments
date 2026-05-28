package main

import (
	"fmt"
	"sync"
)

func maindd() {
	ids := []string{"1", "2", "3", "4", "5"}

	orderChan := make(chan order)
	var wg sync.WaitGroup

	go receiveOrders(orderChan, ids...)

	wg.Add(1)
	go func() {
		defer wg.Done()
		validateOrders(orderChan)
	}()

	wg.Wait()
}

func receiveOrders(out chan<- order, ids ...string) {
	price := -2

	for _, id := range ids {
		out <- order{price: price, id: id}
		price += 1
	}
	close(out)
}

func validateOrders(in <-chan order) {
	for order := range in {
		if order.price >= 0 {
			fmt.Printf("Valid Order: %+v\n", order)
		} else {
			fmt.Printf("Invalid Order dropped: %s\n", order.id)
		}
	}
}
