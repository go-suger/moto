package moto

import "context"

type StateMachineBuilder[S, E comparable, C any] struct {
}

func New[S, E comparable, C any]() *StateMachineBuilder[S, E, C] {
	return &StateMachineBuilder[S, E, C]{}
}

func (b *StateMachineBuilder[S, E, C]) ExternalTransition() *Transition[S, E, C] {
	return &Transition[S, E, C]{}
}

func (b *StateMachineBuilder[S, E, C]) Build(ctx context.Context) (*StateMachine[S, E, C], error) {
	return &StateMachine[S, E, C]{}, nil
}
