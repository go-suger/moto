package moto

type Condition[C any] interface {
	isSatisfied(context C) bool
}

type SimpleConditionImpl[C any] struct {
	condition func(context C) bool
}

func (sc *SimpleConditionImpl[C]) isSatisfied(context C) bool {
	return sc.condition(context)
}
