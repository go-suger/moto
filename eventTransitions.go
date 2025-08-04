package moto

import (
	"fmt"
	"sync"
)

// EventTransitions 实现
type EventTransitions[S, E comparable, C any] struct {
	mu               sync.RWMutex
	eventTransitions map[E][]*Transition[S, E, C]
}

func newEventTransitions[S, E comparable, C any]() *EventTransitions[S, E, C] {
	return &EventTransitions[S, E, C]{
		eventTransitions: make(map[E][]*Transition[S, E, C]),
	}
}

// Put 添加一个 transition
func (et *EventTransitions[S, E, C]) Put(event E, transition *Transition[S, E, C]) error {
	et.mu.Lock()
	defer et.mu.Unlock()

	transitions, ok := et.eventTransitions[event]
	if !ok {
		et.eventTransitions[event] = []*Transition[S, E, C]{transition}
		return nil
	}
	if err := et.verify(transitions, transition); err != nil {
		return err
	}
	et.eventTransitions[event] = append(transitions, transition)
	return nil
}

// Get 获取某个 event 的所有 transition
func (et *EventTransitions[S, E, C]) Get(event E) []*Transition[S, E, C] {
	et.mu.RLock()
	defer et.mu.RUnlock()
	return et.eventTransitions[event]
}

// AllTransitions 获取所有 transition
func (et *EventTransitions[S, E, C]) AllTransitions() []*Transition[S, E, C] {
	et.mu.RLock()
	defer et.mu.RUnlock()
	var all []*Transition[S, E, C]
	for _, transitions := range et.eventTransitions {
		all = append(all, transitions...)
	}
	return all
}

// verify 保证同一 source/target/event 只允许一个 transition
func (et *EventTransitions[S, E, C]) verify(existing []*Transition[S, E, C], newT *Transition[S, E, C]) error {
	for _, t := range existing {
		if t.source == newT.source && t.target == newT.target && t.event == newT.event {
			return fmt.Errorf("transition %+v already exists, you cannot add another one", t)
		}
	}
	return nil
}
