package moto

import (
	"errors"
)

type Transition[S, E comparable, C any] struct {
	source    *State[S, E, C]
	target    *State[S, E, C]
	event     E
	condition Condition[C]
	action    Action[S, E, C]
}

func NewTransition[S, E comparable, C any](source *State[S, E, C], target *State[S, E, C], event E, condition Condition[C], action Action[S, E, C]) *Transition[S, E, C] {
	return &Transition[S, E, C]{
		source:    source,
		target:    target,
		event:     event,
		condition: condition,
		action:    action,
	}
}

func (tr Transition[S, E, C]) transit(ctx *C) (state *State[S, E, C], err error) {
	if tr.source == tr.target {
		err = errors.New("source and target states cannot be the same")
		return
	}

	if tr.condition != nil && tr.condition.isSatisfied(*ctx) {
		if tr.action != nil {
			if err = tr.action.execute(tr.source.State(), tr.target.State(), tr.event, ctx); err != nil {
				return
			} else {
				return tr.target, nil
			}
		} else {
			return tr.target, nil
		}
	}

	return tr.source, nil
}
