package xctx_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jussi-kalliokoski/xctx"
)

func TestWithValue(t *testing.T) {
	t.Run("value is set", func(t *testing.T) {
		type Value struct{ Foo string }
		var ValueCtx = xctx.WithValue[Value]()
		expectedValue := Value{Foo: "abcd"}
		expectedOK := true
		ctx := ValueCtx.Context(context.Background(), expectedValue)

		receivedValue, receivedOK := ValueCtx.Get(ctx)

		if expectedOK != receivedOK {
			t.Fatalf("expected %#v, received %#v", expectedOK, receivedOK)
		}
		if expectedValue != receivedValue {
			t.Fatalf("expected %#v, received %#v", expectedValue, receivedValue)
		}
	})

	t.Run("value is not set", func(t *testing.T) {
		type Value struct{ Foo string }
		var ValueCtx = xctx.WithValue[Value]()
		var expectedValue Value
		var expectedOK bool
		ctx := context.Background()

		receivedValue, receivedOK := ValueCtx.Get(ctx)

		if expectedOK != receivedOK {
			t.Fatalf("expected %#v, received %#v", expectedOK, receivedOK)
		}
		if expectedValue != receivedValue {
			t.Fatalf("expected %#v, received %#v", expectedValue, receivedValue)
		}
	})
}

func ExampleWithValue() {
	// Initialization stage
	type Value struct{ Foo string }
	var ValueCtx = xctx.WithValue[Value]()

	// Set value
	ctx := ValueCtx.Context(context.Background(), Value{Foo: "abcd"})

	// Get value
	value, ok := ValueCtx.Get(ctx)

	// Output: abcd true
	fmt.Println(value.Foo, ok)
}
