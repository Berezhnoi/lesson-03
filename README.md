# Lesson 03 Go Project

A small Go project covering three focused practice areas:

- payment logic
- wallet pointer/value receiver correctness
- dependency injection and testing for orders

## Project structure

- `payment/` — payment-related logic and tests
- `wallet/` — wallet behavior and pointer/value receiver fix
- `orders/` — `OrderService`, `OrderStore` abstraction, and tests
- `prompts/` — prompt examples for AI-assisted refactoring
- `REPORT.md` — summary/report for the homework

## Modules overview

### `payment`
Contains payment domain logic and tests for payment processing.

### `wallet`
Covers the common Go bug where a slice of values is mutated through a copy during range iteration. The fix ensures state is updated in the real slice elements.

### `orders`
Implements a simple order service that depends on a minimal `OrderStore` interface instead of a concrete database object. This keeps the code testable and loosely coupled.

## Run locally

From the project root:

```bash
go test ./...
```

Run each package separately:

```bash
go test ./payment/...
go test ./wallet/...
go test ./orders/...
```

## Format Go files

```bash
gofmt -w .
```

## Notes

- No real database is used in tests for the orders package.
- The service uses dependency injection via `NewOrderService(store)`.
- The project is intentionally simple and focused on clean Go patterns.
