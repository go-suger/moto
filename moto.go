package moto

type StateMachine[S, E comparable, C any] struct {
}

func (sm *StateMachine[S, E, C]) FireEvent(sourceState S, event E, ctx C) error {
	return nil
}

func (sm *StateMachine[S, E, C]) GenerateMermaidGraph() string {
	return "nil"
}
