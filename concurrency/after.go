package main

import (
	"fmt"
	"sync"
)

func mainskl(jkj) {
	orderCh := make(chan Order, 10)

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			order := createOrder(uint(i))

			orderCh <- order
		}(i)
	}

	go func() {
		wg.Wait()
		close(orderCh)
	}()

	for order := range orderCh {
		fmt.Println(order)
	}
}

func createOrder(orderID uint) Order {
	var wg sync.WaitGroup

	var vendor string
	var delivery string

	wg.Add(2)

	go func() {
		defer wg.Done()
		vendor = getVendorInfo()
	}()

	go func() {
		defer wg.Done()
		delivery = getDeliveryInfo()
	}()

	wg.Wait()

	return Order{
		ID:       orderID,
		Vendor:   vendor,
		Delivery: delivery,
	}
}

func getVendorInfo() string {
	return "Shahrvand"
}

func getDeliveryInfo() string {
	return "Snapp"
}
