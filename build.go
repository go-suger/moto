package moto

import (
	"errors"
)

type StateMachineBuilder[S, E comparable, C any] struct {
	stateMap     map[S]*State[S, E, C]
	stateMachine *StateMachine[S, E, C]

	errs []error
}

func New[S, E comparable, C any]() *StateMachineBuilder[S, E, C] {
	stateMap := make(map[S]*State[S, E, C])
	return &StateMachineBuilder[S, E, C]{
		stateMap:     stateMap,
		stateMachine: newStateMachine[S, E, C](stateMap),
		errs:         make([]error, 0),
	}
}

func (b *StateMachineBuilder[S, E, C]) ExternalTransition() Form[S, E, C] {
	return newTransitionBuild(b.stateMap, b.errs)
}

func (b *StateMachineBuilder[S, E, C]) Build() (*StateMachine[S, E, C], error) {
	if len(b.errs) > 0 {
		return nil, errors.Join(b.errs...)
	}

	return b.stateMachine, nil
}
