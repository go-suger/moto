package moto

type StateMachineBuilder[S, E comparable, C any] struct {
	stateMap     map[S]*State[S, E, C]
	stateMachine *StateMachine[S, E, C]
}

func New[S, E comparable, C any]() *StateMachineBuilder[S, E, C] {
	stateMap := make(map[S]*State[S, E, C])
	return &StateMachineBuilder[S, E, C]{
		stateMap:     stateMap,
		stateMachine: newStateMachine[S, E, C](stateMap),
	}
}

func (b *StateMachineBuilder[S, E, C]) ExternalTransition() *TransitionBuild[S, E, C] {
	return NewTransitionBuild(b.stateMap)
}

func (b *StateMachineBuilder[S, E, C]) Build() (*StateMachine[S, E, C], error) {

	return b.stateMachine, nil
}
