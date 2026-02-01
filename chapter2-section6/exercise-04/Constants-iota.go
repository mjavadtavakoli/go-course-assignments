 
package main

import "fmt"

const (
	OrderPending = iota
	OrderConfirmed
	OrderShipped
	OrderDelivered
	OrderCancelled
)

func statusToString(status int) string {
	switch status {
	case OrderPending:
		return "Pending"
	case OrderConfirmed:
		return "Confirmed"
	case OrderShipped:
		return "Shipped"
	case OrderDelivered:
		return "Delivered"
	case OrderCancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("Status:", statusToString(OrderDelivered))
	fmt.Println("Status code:", OrderDelivered)
}


//response : 

/*Status: Delivered
Status code: 3
*/
