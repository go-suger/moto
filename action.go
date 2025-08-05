package moto

import (
	"context"
)

type Action[S, E comparable, C any] interface {
	execute(ctx context.Context, from, to S, event E, context *C) error
}

type SimpleActionImpl[S, E comparable, C any] struct {
	action func(ctx context.Context, from, to S, event E, context *C) error
}

func (sa *SimpleActionImpl[S, E, C]) execute(ctx context.Context, from, to S, event E, context *C) error {
	return sa.action(ctx, from, to, event, context)
}
