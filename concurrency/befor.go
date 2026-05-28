package main

import (
	"fmt"
	"sync"
)

func main() {
	ids := []string{"1", "2", "3", "4", "5"}

	orderChan := make(chan order)
	deliveryChan := make(chan order)

	var wg sync.WaitGroup

	go receiveOrders(orderChan, ids...)

	wg.Add(1)
	go func() {
		defer wg.Done()
		validateOrders(orderChan, deliveryChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		processDelivery(deliveryChan)
	}()

	wg.Wait()
	fmt.Println("All orders processed.")
}

func receiveOrders(out chan<- order, ids ...string) {
	price := -2
	for _, id := range ids {
		out <- order{price: price, id: id}
		price += 1
	}
	close(out)
}

func validateOrders(in <-chan order, validOrdersOut chan<- order) {
	for order := range in {
		if order.price >= 0 {
			fmt.Printf("Validated Order: %+v\n", order)
			validOrdersOut <- order
		} else {
			fmt.Printf("Invalid Order dropped: %s (price: %d)\n", order.id, order.price)
		}
	}
	close(validOrdersOut)
}

func processDelivery(in <-chan order) {
	for order := range in {
		fmt.Printf("--- Processing Delivery for Order ID: %s ---\n", order.id)
	}
}
