package moto

import (
	"context"
	"errors"
)

type StateMachine[S, E comparable, C any] struct {
	stateMap map[S]*State[S, E, C]
}

func newStateMachine[S, E comparable, C any](stateMap map[S]*State[S, E, C]) *StateMachine[S, E, C] {
	return &StateMachine[S, E, C]{stateMap: stateMap}
}

func (sm *StateMachine[S, E, C]) FireEvent(ctx context.Context, sourceState S, event E, context *C) (stateVal S, err error) {
	if context == nil {
		err = errors.New("ctx is nil")
		return
	}

	transition, err := sm.routeTransition(ctx, sourceState, event, *context)
	if err != nil {
		return
	}

	if transition == nil {
		err = errors.New("transition is nil")
		return
	}

	state, err := transition.transit(ctx, context)
	if err != nil {
		return
	}

	return state.State(), nil
}

func (sm *StateMachine[S, E, C]) routeTransition(ctx context.Context, sourceStateVal S, event E, context C) (*Transition[S, E, C], error) {
	sourceState, err := sm.getState(sourceStateVal)
	if err != nil {
		return nil, err
	}

	var transition = sourceState.getEventTransitions(event)
	if transition == nil {
		return nil, errors.New("transition is nil")
	}

	var transit *Transition[S, E, C] = nil

	if transition.condition == nil {
		transit = transition
	} else if transition.condition.isSatisfied(ctx, context) {
		transit = transition
	}

	return transit, nil
}

func (sm *StateMachine[S, E, C]) getState(currentState S) (state *State[S, E, C], err error) {
	state, ok := sm.stateMap[currentState]
	if !ok {
		err = errors.New("state not found")
		return
	}

	return state, nil
}

func (sm *StateMachine[S, E, C]) GenerateMermaidGraph() string {
	return "nil"
}

type State[S, E comparable, C any] struct {
	state            S
	eventTransitions *EventTransitions[S, E, C]
}

func newState[S, E comparable, C any](state S) *State[S, E, C] {
	return &State[S, E, C]{
		state:            state,
		eventTransitions: newEventTransitions[S, E, C](),
	}
}

func (s *State[S, E, C]) State() S {
	return s.state
}

func (s *State[S, E, C]) getEventTransitions(event E) *Transition[S, E, C] {
	return s.eventTransitions.Get(event)
}

func (s *State[S, E, C]) addTransition(event E, target *State[S, E, C]) (*Transition[S, E, C], error) {
	transition := newTransition(s, target, event, nil, nil)

	if err := s.eventTransitions.Put(event, transition); err != nil {
		return nil, errors.New("failed to add event to transition")
	}

	return transition, nil
}
