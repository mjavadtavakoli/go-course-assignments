## Mini OMS – In-Memory Order & Voucher Engine

This package implements a small, in-memory Order Management System (OMS) as described in the assignment.

### Pointer vs Value Decisions

- **Order aggregate**: `Order` is passed and stored as a pointer (`*Order`) in the repository so that state changes (e.g. `Pay`, `Cancel`, `ChangeState`) are performed on a single canonical instance. This mirrors typical domain-driven design aggregates.
- **Repository map**: The in-memory repository keeps a `map[int]*Order` internally. This avoids copying large aggregates on each save and makes updates inexpensive.
- **Read paths**: `FindByID` returns a *copy* of the order (`*Order` whose underlying value has been copied). The `Items` slice is cloned via `SnapshotItems` to avoid callers mutating the repository’s internal slice storage. `List` returns a `[]Order` (by value, as required) built from clones for the same reason.
- **Order items and products**: `Product` and `OrderItem` are small immutable-ish structs and are passed/returned by value.

### Interface Reasoning

- **Voucher interface**: `Voucher` abstracts discount behavior:
  - `PercentageVoucher` and `FixedAmountVoucher` are two concrete strategies.
  - The `Order` knows only about `Voucher`, not which implementation is used. This keeps order logic closed for modification when adding new voucher types.
  - The repository and higher-order functions do not depend on concrete voucher types.
- **OrderRepository interface**: Defines an in-memory repository contract (`Save`, `FindByID`, `List`, `Delete`, `Clear`), allowing:
  - Easy swapping of persistence (e.g., to a database) without changing domain logic.
  - Tests to depend on the interface while using `InMemoryOrderRepo` as the implementation.

### State Validation Logic

- The `OrderState` is an `int`-backed enum:
  - `Created`, `Paid`, `VendorAccepted`, `Shipped`, `Delivered`, `Cancelled`.
- `ChangeState` uses a `switch` on the current state, and nested switches on the `newState`:
  - **Created →** `Paid`, `Cancelled`
  - **Paid →** `VendorAccepted`, `Cancelled`
  - **VendorAccepted →** `Shipped`, `Cancelled`
  - **Shipped →** `Delivered`, `Cancelled`
  - **Delivered** is terminal (no outgoing transitions).
  - **Cancelled** is terminal (any further transition returns an `InvalidStateTransitionError`).
- Any disallowed transition returns `InvalidStateTransitionError{From, To}` with a descriptive message.
- Higher-level methods (`Pay`, `Cancel`) delegate to `ChangeState` after enforcing business invariants (e.g., non-empty items, positive total).

### Voucher Validation and Nil Handling

- **Percentage vouchers**:
  - `NewPercentageVoucher` validates `percent` in \[1, 100\] and non-empty `code`.
  - `Apply` computes `discount = total * percent / 100` and caps the result at 0 (never negative).
- **Fixed amount vouchers**:
  - `NewFixedAmountVoucher` requires `amount > 0` and non-empty `code`.
  - `Apply` subtracts `amount` and caps at 0 (never negative).
- Both `Apply` methods:
  - Gracefully handle `total <= 0` by returning the input unchanged.
  - Are pointer receivers but guard against `nil` receivers: if `v == nil`, `Apply` returns the original total and `Code` returns `""`.
- **Nil interface usage**:
  - `Order.ApplyVoucher` checks `v == nil` before assigning to prevent storing a nil `Voucher` and to avoid panics on later calls to `Apply`.
  - `TotalAmount` checks `o.Voucher != nil` before calling `Apply`, which correctly handles the zero `Voucher` value.

### Business Rules and Edge Cases

- **State rules**:
  - Cannot ship before paid or vendor-accepted: enforced by the explicit state-transition table in `ChangeState`.
  - Cannot deliver before shipped: only `Shipped → Delivered` is allowed.
  - Cannot transition after cancellation or delivery: both are treated as terminal states.
- **Totals and items**:
  - `NewOrderItem` rejects zero or negative quantities and negative prices.
  - `TotalAmount` returns an error if:
    - The order has no items.
    - Any item has invalid quantity or price.
  - A `defer` in `TotalAmount` guarantees that on any success path, the total is strictly greater than zero; otherwise it mutates the named error return.
  - `Pay` calls `TotalAmount` and rejects any non-positive total, ensuring you cannot pay for an empty or fully-discounted order.
- **Vouchers**:
  - `ApplyVoucher`:
    - Prevents applying more than one voucher (`o.Voucher != nil`).
    - Requires the order to be in the `Created` state (no voucher changes after the payment flow begins).
  - Voucher `Apply` methods never produce negative totals (both clamp to 0).

### Repository Behavior

- `InMemoryOrderRepo`:
  - Initializes its map in `NewInMemoryOrderRepo` and lazily in `Save` if needed.
  - Uses comma-ok (`order, ok := r.orders[id]`) in `FindByID` and `Delete` to detect missing keys and return descriptive errors.
  - `Clear` iterates and deletes keys, avoiding exposing the internal map or replacing it with `nil`.

### Higher-Order Function and Closures

- `FilterOrders`:
  - Allocates a new slice and appends matching orders by evaluating `fn(Order) bool`.
  - This ensures the original slice is never modified.
- `MinAmountFilter`:
  - Demonstrates a **closure**:
    - Captures `min` in an outer function.
    - Returns `func(Order) bool` that calls `TotalAmount` and compares against `min`.
  - Example usage:

  ```go
  allOrders := repo.List()
  expensiveOnly := FilterOrders(allOrders, MinAmountFilter(10_000))
  paidOrders := FilterOrders(allOrders, func(o Order) bool {
      return o.State == Paid
  })
  ```

