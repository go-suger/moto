package moto

type TransitionBuild[S, E comparable, C any] struct {
}

func (tt *TransitionBuild[S, E, C]) Form(state ...S) *TransitionBuild[S, E, C] {

	return tt
}

func (tt *TransitionBuild[S, E, C]) To(state S) *TransitionBuild[S, E, C] {

	return tt
}

func (tt *TransitionBuild[S, E, C]) On(event E) *TransitionBuild[S, E, C] {

	return tt
}

func (tt *TransitionBuild[S, E, C]) When(condition Condition[C]) *TransitionBuild[S, E, C] {

	return tt
}

func (tt *TransitionBuild[S, E, C]) WhenFunc(condition func(context C) bool) *TransitionBuild[S, E, C] {
	tt.When(&SimpleConditionImpl[C]{condition: condition})
	return tt
}

func (tt *TransitionBuild[S, E, C]) Perform(action Action[S, E, C]) *TransitionBuild[S, E, C] {

	return tt
}

func (tt *TransitionBuild[S, E, C]) PerformFunc(action func(from, to S, event E, context *C) error) *TransitionBuild[S, E, C] {
	tt.Perform(&SimpleActionImpl[S, E, C]{action: action})
	return tt
}

func (tt *TransitionBuild[S, E, C]) validation() error {

	return nil
}
