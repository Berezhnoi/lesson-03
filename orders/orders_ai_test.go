package orders

import (
	"errors"
	"reflect"
	"testing"
)

// fakeOrderStore is a compact AI-generated test double for OrderStore.
// It records the exact Exec call so the service can be validated without a real DB.
type fakeOrderStore struct {
	query string
	args  []any
	err   error
	calls int
}

func (f *fakeOrderStore) Exec(query string, args ...any) error {
	f.calls++
	f.query = query
	f.args = append([]any(nil), args...)
	return f.err
}

func TestOrderServicePlaceOrder_AI(t *testing.T) {
	const (
		query = "INSERT INTO orders (id, amount) VALUES (?, ?)"
		id    = "order-ai"
		amt   = 77.5
	)

	t.Run("success path", func(t *testing.T) {
		store := &fakeOrderStore{}
		service := NewOrderService(store)

		if err := service.PlaceOrder(id, amt); err != nil {
			t.Fatalf("PlaceOrder() returned unexpected error: %v", err)
		}
		if store.calls != 1 {
			t.Fatalf("Exec() call count = %d, want 1", store.calls)
		}
		if store.query != query {
			t.Fatalf("Exec() query = %q, want %q", store.query, query)
		}
		wantArgs := []any{id, amt}
		if !reflect.DeepEqual(store.args, wantArgs) {
			t.Fatalf("Exec() args = %#v, want %#v", store.args, wantArgs)
		}
	})

	t.Run("store error is propagated", func(t *testing.T) {
		wantErr := errors.New("database unavailable")
		store := &fakeOrderStore{err: wantErr}
		service := NewOrderService(store)

		err := service.PlaceOrder(id, amt)
		if !errors.Is(err, wantErr) {
			t.Fatalf("PlaceOrder() error = %v, want %v", err, wantErr)
		}
	})
}
