package moto

type StateMachineBuilder[S, E comparable, C any] struct {
}

func New[S, E comparable, C any]() *StateMachineBuilder[S, E, C] {
	return &StateMachineBuilder[S, E, C]{}
}

func (b *StateMachineBuilder[S, E, C]) ExternalTransition() *TransitionBuild[S, E, C] {
	return &TransitionBuild[S, E, C]{}
}

func (b *StateMachineBuilder[S, E, C]) Build() (*StateMachine[S, E, C], error) {
	stateMachine := NewStateMachine[S, E, C]()

	return stateMachine, nil
}
