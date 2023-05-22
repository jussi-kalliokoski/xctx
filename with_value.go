package xctx

import (
	"context"
	"sync/atomic"
)

type withValueKey uint64

var currentWithValueKey uint64

// WithValueContextBuilder is a utility for adding typed values into context.
type WithValueContextBuilder[V any] struct {
	key withValueKey
}

// WithValue is a utility for adding typed values into context, removing
// some boilerplate (creating a custom key type and casting the value into the
// correct type) that can be error-prone as context values are not typed in the
// standard library.
func WithValue[V any]() *WithValueContextBuilder[V] {
	key := withValueKey(atomic.AddUint64(&currentWithValueKey, 1))
	return &WithValueContextBuilder[V]{key: key}
}

// Context returns a context containing the given value.
func (b *WithValueContextBuilder[V]) Context(ctx context.Context, v V) context.Context {
	return context.WithValue(ctx, b.key, v)
}

// Get returns the value stored in the context and true, or a zero-value of
// the value type and false if the value is not present in the context.
func (b *WithValueContextBuilder[V]) Get(ctx context.Context) (value V, ok bool) {
	value, ok = ctx.Value(b.key).(V)
	return value, ok
}
