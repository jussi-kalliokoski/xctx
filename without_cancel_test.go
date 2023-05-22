package xctx_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jussi-kalliokoski/xctx"
)

func TestWithoutCancel(t *testing.T) {
	t.Run("Deadline should be empty", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		ctx = xctx.WithoutCancel(ctx)
		var expectedDeadline time.Time
		var expectedOK bool

		receivedDeadline, receivedOK := ctx.Deadline()

		if expectedOK != receivedOK {
			t.Fatalf("expected %#v, received %#v", expectedOK, receivedOK)
		}
		if expectedDeadline != receivedDeadline {
			t.Fatalf("expected %#v, received %#v", expectedDeadline, receivedDeadline)
		}
	})

	t.Run("Done should be empty", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ctx = xctx.WithoutCancel(ctx)
		cancel()
		var expected <-chan struct{}

		received := ctx.Done()

		if expected != received {
			t.Fatalf("expected %#v, received %#v", expected, received)
		}
	})

	t.Run("Err should be empty", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ctx = xctx.WithoutCancel(ctx)
		cancel()
		var expected error

		received := ctx.Err()

		if expected != received {
			t.Fatalf("expected %#v, received %#v", expected, received)
		}
	})

	t.Run("Value should be propagated", func(t *testing.T) {
		type Key struct{}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		expected := 1234
		ctx = context.WithValue(ctx, Key{}, expected)
		ctx = xctx.WithoutCancel(ctx)
		defer cancel()

		received := ctx.Value(Key{})

		if expected != received {
			t.Fatalf("expected %#v, received %#v", expected, received)
		}
	})
}

func ExampleWithoutCancel() {
	type Key struct{}
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, Key{}, 1234)
	ctx = xctx.WithoutCancel(ctx)
	cancel()

	doneCh := ctx.Done()
	err := ctx.Err()
	val := ctx.Value(Key{})

	// Output: <nil> <nil> 1234
	fmt.Println(doneCh, err, val)
}
