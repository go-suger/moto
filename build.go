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
		stateMachine: newStateMachine(stateMap),
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

type TransitionBuild[S, E comparable, C any] struct {
	stateMap map[S]*State[S, E, C]
	errs     []error

	sources     []*State[S, E, C]
	target      *State[S, E, C]
	transitions []*Transition[S, E, C]
}

func newTransitionBuild[S, E comparable, C any](stateMap map[S]*State[S, E, C], errs []error) *TransitionBuild[S, E, C] {
	return &TransitionBuild[S, E, C]{
		stateMap:    stateMap,
		sources:     make([]*State[S, E, C], 0),
		transitions: make([]*Transition[S, E, C], 0),
		errs:        errs,
	}
}

type Form[S, E comparable, C any] interface {
	Form(stateVals ...S) To[S, E, C]
}

func (tt *TransitionBuild[S, E, C]) Form(stateVals ...S) To[S, E, C] {
	for _, stateVal := range stateVals {
		tt.sources = append(tt.sources, getState(tt.stateMap, stateVal))
	}

	return tt
}

type To[S, E comparable, C any] interface {
	To(state S) On[S, E, C]
}

func (tt *TransitionBuild[S, E, C]) To(state S) On[S, E, C] {
	tt.target = getState(tt.stateMap, state)

	return tt
}

type On[S, E comparable, C any] interface {
	On(event E) When[S, E, C]
}

// On a single event can only have a single Transition
func (tt *TransitionBuild[S, E, C]) On(event E) When[S, E, C] {
	for _, source := range tt.sources {
		transition, err := source.addTransition(event, tt.target)
		if err != nil {
			tt.errs = append(tt.errs, err)
			return tt
		}

		tt.transitions = append(tt.transitions, transition)
	}

	return tt
}

type When[S, E comparable, C any] interface {
	When(condition Condition[C]) Perform[S, E, C]
	WhenFunc(conditionFunc func(context C) bool) Perform[S, E, C]
}

func (tt *TransitionBuild[S, E, C]) When(condition Condition[C]) Perform[S, E, C] {
	for _, transition := range tt.transitions {
		transition.condition = condition
	}

	return tt
}

func (tt *TransitionBuild[S, E, C]) WhenFunc(conditionFunc func(context C) bool) Perform[S, E, C] {
	tt.When(&SimpleConditionImpl[C]{condition: conditionFunc})
	return tt
}

type Perform[S, E comparable, C any] interface {
	Perform(action Action[S, E, C])
	PerformFunc(actionFunc func(from, to S, event E, context *C) error)
}

func (tt *TransitionBuild[S, E, C]) Perform(action Action[S, E, C]) {
	for _, transition := range tt.transitions {
		transition.action = action
	}
}

func (tt *TransitionBuild[S, E, C]) PerformFunc(actionFunc func(from, to S, event E, context *C) error) {
	tt.Perform(&SimpleActionImpl[S, E, C]{action: actionFunc})
}
