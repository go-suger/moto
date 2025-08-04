package moto

import (
	"errors"
)

type StateMachine[S, E comparable, C any] struct {
	// 状态map
	stateMap map[S]*State[S, E, C]
}

func newStateMachine[S, E comparable, C any](stateMap map[S]*State[S, E, C]) *StateMachine[S, E, C] {
	return &StateMachine[S, E, C]{stateMap: stateMap}
}

func (sm *StateMachine[S, E, C]) FireEvent(sourceState S, event E, ctx *C) (stateVal S, err error) {
	if ctx == nil {
		err = errors.New("ctx is nil")
		return
	}

	transition, err := sm.routeTransition(sourceState, event, ctx)
	if err != nil {
		return
	}

	if transition == nil {
		err = errors.New("transition is nil")
		return
	}

	state, err := transition.transit(ctx)
	if err != nil {
		return
	}

	return state.State(), nil
}

func (sm *StateMachine[S, E, C]) routeTransition(sourceStateVal S, event E, ctx *C) (*Transition[S, E, C], error) {
	sourceState, err := sm.getState(sourceStateVal)
	if err != nil {
		return nil, err
	}

	var transitions = sourceState.getEventTransitions(event)

	if transitions == nil || len(transitions) == 0 {
		return nil, errors.New("transition is nil")
	}

	var transit *Transition[S, E, C] = nil
	for _, transition := range transitions {
		if transition.condition == nil {
			transit = transition
		} else if transition.condition.isSatisfied(*ctx) {
			transit = transition
			break
		}
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

func NewState[S, E comparable, C any](state S) *State[S, E, C] {
	return &State[S, E, C]{
		state:            state,
		eventTransitions: NewEventTransitions[S, E, C](),
	}
}

func (s *State[S, E, C]) State() S {
	return s.state
}

func (s *State[S, E, C]) getEventTransitions(event E) []*Transition[S, E, C] {
	return s.eventTransitions.Get(event)
}

func (s *State[S, E, C]) addTransition(event E, target *State[S, E, C]) (*Transition[S, E, C], error) {
	transition := NewTransition(s, target, event, nil, nil)
	err := s.eventTransitions.Put(event, transition)
	if err != nil {
		return nil, errors.New("failed to add event to transition")
	}
	return transition, nil
}
