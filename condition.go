package moto

import (
	"context"
)

type Condition[C any] interface {
	isSatisfied(ctx context.Context, context C) bool
}

type SimpleConditionImpl[C any] struct {
	condition func(ctx context.Context, context C) bool
}

func (sc *SimpleConditionImpl[C]) isSatisfied(ctx context.Context, context C) bool {
	return sc.condition(ctx, context)
}
