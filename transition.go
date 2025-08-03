package moto

type Transition[S, E comparable, C any] struct {
}

func (tt *Transition[S, E, C]) Form(state ...S) *Transition[S, E, C] {

	return tt
}

func (tt *Transition[S, E, C]) To(state S) *Transition[S, E, C] {

	return tt
}

func (tt *Transition[S, E, C]) On(event E) *Transition[S, E, C] {

	return tt
}

func (tt *Transition[S, E, C]) When(condition func(form, to S, context C)) *Transition[S, E, C] {

	return tt
}

func (tt *Transition[S, E, C]) Perform(action func(form, to S, context C)) *Transition[S, E, C] {

	return tt
}

func (tt *Transition[S, E, C]) validation() error {

	return nil
}
