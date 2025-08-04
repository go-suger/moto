package moto

func getState[S, E comparable, C any](stateMap map[S]*State[S, E, C], stateVal S) *State[S, E, C] {
	newState, ok := stateMap[stateVal]
	if !ok {
		newState = NewState[S, E, C](stateVal)
		stateMap[stateVal] = newState
	}
	return newState
}

type TransitionBuild[S, E comparable, C any] struct {
	stateMap    map[S]*State[S, E, C]
	sources     []*State[S, E, C]
	target      *State[S, E, C]
	transitions []*Transition[S, E, C]
}

func NewTransitionBuild[S, E comparable, C any](stateMap map[S]*State[S, E, C]) *TransitionBuild[S, E, C] {
	return &TransitionBuild[S, E, C]{
		stateMap:    stateMap,
		sources:     make([]*State[S, E, C], 0),
		transitions: make([]*Transition[S, E, C], 0),
	}
}

func (tt *TransitionBuild[S, E, C]) Form(stateVals ...S) *TransitionBuild[S, E, C] {
	for _, stateVal := range stateVals {
		tt.sources = append(tt.sources, getState(tt.stateMap, stateVal))
	}

	return tt
}

func (tt *TransitionBuild[S, E, C]) To(state S) *TransitionBuild[S, E, C] {
	tt.target = getState(tt.stateMap, state)
	return tt
}

func (tt *TransitionBuild[S, E, C]) On(event E) *TransitionBuild[S, E, C] {
	for _, source := range tt.sources {
		transition, _ := source.addTransition(event, tt.target)
		tt.transitions = append(tt.transitions, transition)
	}

	return tt
}

func (tt *TransitionBuild[S, E, C]) When(condition Condition[C]) *TransitionBuild[S, E, C] {
	for _, transition := range tt.transitions {
		transition.condition = condition
	}

	return tt
}

func (tt *TransitionBuild[S, E, C]) Perform(action Action[S, E, C]) *TransitionBuild[S, E, C] {
	for _, transition := range tt.transitions {
		transition.action = action
	}

	return tt
}

func (tt *TransitionBuild[S, E, C]) WhenFunc(condition func(context C) bool) *TransitionBuild[S, E, C] {
	tt.When(&SimpleConditionImpl[C]{condition: condition})
	return tt
}

func (tt *TransitionBuild[S, E, C]) PerformFunc(action func(from, to S, event E, context *C) error) *TransitionBuild[S, E, C] {
	tt.Perform(&SimpleActionImpl[S, E, C]{action: action})
	return tt
}

func (tt *TransitionBuild[S, E, C]) validation() error {

	return nil
}
