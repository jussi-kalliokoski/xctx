package xctx

import (
	"context"
	"time"
)

// WithoutCancel returns a context whose cancellation is not propagated, i.e.
// Deadline, Done, and Err all return zero-values, but the Value calls are
// propagated to the parent context.
func WithoutCancel(ctx context.Context) context.Context {
	return ctxWithoutCancel{parent: ctx}
}

type ctxWithoutCancel struct {
	parent context.Context
}

func (c ctxWithoutCancel) Deadline() (deadline time.Time, ok bool) {
	return deadline, ok
}

func (c ctxWithoutCancel) Done() <-chan struct{} {
	return nil
}

func (c ctxWithoutCancel) Err() error {
	return nil
}

func (c ctxWithoutCancel) Value(key any) any {
	return c.parent.Value(key)
}
