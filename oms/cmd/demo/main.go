package main

import (
	"fmt"

	"oms"
)

func main() {
	repo := oms.NewInMemoryOrderRepo()

	o := oms.NewOrder(1)
	_ = o.AddItem(oms.Product{ID: 10, Name: "Keyboard", Price: 5000}, 1)
	_ = o.AddItem(oms.Product{ID: 11, Name: "Mouse", Price: 2000}, 2)

	v, _ := oms.NewPercentageVoucher("OFF10", 10)
	_ = o.ApplyVoucher(v)

	total, err := o.TotalAmount()
	fmt.Println("total:", total, "err:", err)

	_ = o.Pay()
	_ = repo.Save(o)

	all := repo.List()
	paid := oms.FilterOrders(all, func(or oms.Order) bool { return or.State == oms.Paid })
	expensive := oms.FilterOrders(all, oms.MinAmountFilter(6000))

	fmt.Println("paid orders:", len(paid))
	fmt.Println("expensive orders:", len(expensive))
}

