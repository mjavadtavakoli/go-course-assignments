package oms

import (
	"errors"
	"fmt"
)

// OrderState represents the lifecycle state of an order.
type OrderState int

const (
	Created OrderState = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)

// InvalidStateTransitionError is returned when an invalid state change is attempted.
type InvalidStateTransitionError struct {
	From OrderState
	To   OrderState
}

func (e InvalidStateTransitionError) Error() string {
	return fmt.Sprintf("invalid state transition from %d to %d", e.From, e.To)
}

// Product represents a purchasable product.
type Product struct {
	ID    int
	Name  string
	Price int
}

// OrderItem ties a product to a quantity inside an order.
type OrderItem struct {
	Product  Product
	Quantity int
}

// NewOrderItem validates and constructs an OrderItem.
func NewOrderItem(product Product, qty int) (OrderItem, error) {
	if qty <= 0 {
		return OrderItem{}, fmt.Errorf("quantity must be positive, got %d", qty)
	}
	if product.Price < 0 {
		return OrderItem{}, fmt.Errorf("product price cannot be negative, got %d", product.Price)
	}
	return OrderItem{Product: product, Quantity: qty}, nil
}

// TotalPrice returns the total price for this order item.
func (oi OrderItem) TotalPrice() int {
	return oi.Product.Price * oi.Quantity
}

// Voucher represents a discount that can be applied to an order total.
type Voucher interface {
	Apply(total int) int
	Code() string
}

// PercentageVoucher applies a percentage based discount.
type PercentageVoucher struct {
	code    string
	percent int
}

// NewPercentageVoucher validates and constructs a PercentageVoucher.
func NewPercentageVoucher(code string, percent int) (*PercentageVoucher, error) {
	if percent <= 0 || percent > 100 {
		return nil, fmt.Errorf("percent must be between 1 and 100, got %d", percent)
	}
	if code == "" {
		return nil, errors.New("voucher code cannot be empty")
	}
	return &PercentageVoucher{
		code:    code,
		percent: percent,
	}, nil
}

func (v *PercentageVoucher) Apply(total int) int {
	if v == nil {
		return total
	}
	if total <= 0 {
		return total
	}
	discount := total * v.percent / 100
	discounted := total - discount
	if discounted < 0 {
		return 0
	}
	return discounted
}

func (v *PercentageVoucher) Code() string {
	if v == nil {
		return ""
	}
	return v.code
}

// FixedAmountVoucher applies a fixed amount discount.
type FixedAmountVoucher struct {
	code   string
	amount int
}

// NewFixedAmountVoucher validates and constructs a FixedAmountVoucher.
func NewFixedAmountVoucher(code string, amount int) (*FixedAmountVoucher, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be positive, got %d", amount)
	}
	if code == "" {
		return nil, errors.New("voucher code cannot be empty")
	}
	return &FixedAmountVoucher{
		code:   code,
		amount: amount,
	}, nil
}

func (v *FixedAmountVoucher) Apply(total int) int {
	if v == nil {
		return total
	}
	if total <= 0 {
		return total
	}
	discounted := total - v.amount
	if discounted < 0 {
		return 0
	}
	return discounted
}

func (v *FixedAmountVoucher) Code() string {
	if v == nil {
		return ""
	}
	return v.code
}

// Order represents the aggregate root of the order domain.
type Order struct {
	ID      int
	Items   []OrderItem
	Voucher Voucher
	State   OrderState
}

// NewOrder constructs a new order with default state.
func NewOrder(id int) *Order {
	return &Order{
		ID:    id,
		State: Created,
	}
}

// ChangeState validates and updates the order state according to business rules.
func (o *Order) ChangeState(newState OrderState) error {
	if o == nil {
		return errors.New("order is nil")
	}

	if o.State == Cancelled {
		return InvalidStateTransitionError{From: o.State, To: newState}
	}

	switch o.State {
	case Created:
		switch newState {
		case Paid, Cancelled:
			o.State = newState
			return nil
		}
	case Paid:
		switch newState {
		case VendorAccepted, Cancelled:
			o.State = newState
			return nil
		}
	case VendorAccepted:
		switch newState {
		case Shipped, Cancelled:
			o.State = newState
			return nil
		}
	case Shipped:
		switch newState {
		case Delivered, Cancelled:
			o.State = newState
			return nil
		}
	case Delivered:
		// Delivered is terminal.
	}

	return InvalidStateTransitionError{From: o.State, To: newState}
}

// AddItem appends a validated item to the order.
func (o *Order) AddItem(product Product, qty int) error {
	if o == nil {
		return errors.New("order is nil")
	}
	item, err := NewOrderItem(product, qty)
	if err != nil {
		return err
	}
	o.Items = append(o.Items, item)
	return nil
}

// ApplyVoucher attaches a voucher to the order ensuring it is only applied once.
func (o *Order) ApplyVoucher(v Voucher) error {
	if o == nil {
		return errors.New("order is nil")
	}
	if o.State != Created {
		return fmt.Errorf("voucher can only be applied in Created state, current state: %d", o.State)
	}
	if o.Voucher != nil {
		return errors.New("voucher already applied")
	}
	if v == nil {
		return errors.New("voucher is nil")
	}
	o.Voucher = v
	return nil
}

// TotalAmount calculates the final total after applying vouchers.
// Uses defer to ensure the total is never non-positive on successful paths.
func (o *Order) TotalAmount() (total int, err error) {
	if o == nil {
		return 0, errors.New("order is nil")
	}

	defer func() {
		if err == nil && total <= 0 {
			err = errors.New("total amount must be greater than zero")
		}
	}()

	if len(o.Items) == 0 {
		return 0, errors.New("order has no items")
	}

	for _, item := range o.Items {
		if item.Quantity <= 0 {
			return 0, fmt.Errorf("invalid quantity %d in order", item.Quantity)
		}
		if item.Product.Price < 0 {
			return 0, fmt.Errorf("invalid product price %d in order", item.Product.Price)
		}
		total += item.TotalPrice()
	}

	if o.Voucher != nil {
		total = o.Voucher.Apply(total)
	}

	return total, nil
}

// Pay validates and transitions an order into the paid state.
func (o *Order) Pay() error {
	if o == nil {
		return errors.New("order is nil")
	}
	total, err := o.TotalAmount()
	if err != nil {
		return fmt.Errorf("cannot pay for order: %w", err)
	}
	if total <= 0 {
		return errors.New("cannot pay for order with non-positive total")
	}
	return o.ChangeState(Paid)
}

// Cancel transitions an order into the cancelled state.
func (o *Order) Cancel() error {
	if o == nil {
		return errors.New("order is nil")
	}
	return o.ChangeState(Cancelled)
}

// SnapshotItems returns a copy of the items slice to protect internal state.
func (o *Order) SnapshotItems() []OrderItem {
	if o == nil || len(o.Items) == 0 {
		return nil
	}
	cp := make([]OrderItem, len(o.Items))
	copy(cp, o.Items)
	return cp
}

// OrderRepository abstracts storage for orders.
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id int) (*Order, error)
	List() []Order
	Delete(id int) error
	Clear()
}

// InMemoryOrderRepo is an in-memory implementation of OrderRepository.
type InMemoryOrderRepo struct {
	orders map[int]*Order
}

// NewInMemoryOrderRepo constructs a new in-memory repository.
func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int]*Order),
	}
}

// Save inserts or updates an order in the repository.
func (r *InMemoryOrderRepo) Save(order *Order) error {
	if r == nil {
		return errors.New("repository is nil")
	}
	if order == nil {
		return errors.New("order is nil")
	}
	if r.orders == nil {
		r.orders = make(map[int]*Order)
	}
	r.orders[order.ID] = order
	return nil
}

// FindByID returns a copy of the order with the given id.
func (r *InMemoryOrderRepo) FindByID(id int) (*Order, error) {
	if r == nil {
		return nil, errors.New("repository is nil")
	}
	order, ok := r.orders[id]
	if !ok {
		return nil, fmt.Errorf("order with id %d not found", id)
	}
	clone := *order
	clone.Items = order.SnapshotItems()
	return &clone, nil
}

// List returns value copies of all orders in the repository.
func (r *InMemoryOrderRepo) List() []Order {
	if r == nil || len(r.orders) == 0 {
		return nil
	}
	result := make([]Order, 0, len(r.orders))
	for _, o := range r.orders {
		clone := *o
		clone.Items = o.SnapshotItems()
		result = append(result, clone)
	}
	return result
}

// Delete removes an order by id.
func (r *InMemoryOrderRepo) Delete(id int) error {
	if r == nil {
		return errors.New("repository is nil")
	}
	if _, ok := r.orders[id]; !ok {
		return fmt.Errorf("order with id %d not found", id)
	}
	delete(r.orders, id)
	return nil
}

// Clear removes all orders from the repository.
func (r *InMemoryOrderRepo) Clear() {
	if r == nil {
		return
	}
	for k := range r.orders {
		delete(r.orders, k)
	}
}

// FilterOrders returns a new slice containing only orders that satisfy the predicate.
func FilterOrders(orders []Order, fn func(Order) bool) []Order {
	if fn == nil {
		return nil
	}
	result := make([]Order, 0, len(orders))
	for _, o := range orders {
		if fn(o) {
			result = append(result, o)
		}
	}
	return result
}

// MinAmountFilter demonstrates a closure that filters orders by minimum total amount.
func MinAmountFilter(min int) func(Order) bool {
	return func(o Order) bool {
		total, err := o.TotalAmount()
		if err != nil {
			return false
		}
		return total >= min
	}
}
