package moto

type Action[S, E comparable, C any] interface {
	execute(from, to S, event E, context *C) error
}

type SimpleActionImpl[S, E comparable, C any] struct {
	action func(from, to S, event E, context *C) error
}

func (sa *SimpleActionImpl[S, E, C]) execute(from, to S, event E, context *C) error {
	return sa.action(from, to, event, context)
}
