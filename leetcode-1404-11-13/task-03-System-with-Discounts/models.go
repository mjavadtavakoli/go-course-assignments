package order

type OrderItem struct {
	ProductID int
	Quantity  int
	Price     float64
}

type Order struct {
	ID       int
	Items    []OrderItem
	Discount float64
	Total    float64
	Status   string
}

type DiscountRule struct {
	MinAmount       float64
	DiscountPercent float64
	Description     string
}
