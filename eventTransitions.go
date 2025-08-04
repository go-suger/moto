package moto

import (
	"fmt"
	"sync"
)

// EventTransitions 实现
type EventTransitions[S, E comparable, C any] struct {
	mu               sync.RWMutex
	eventTransitions map[E]*Transition[S, E, C]
}

func newEventTransitions[S, E comparable, C any]() *EventTransitions[S, E, C] {
	return &EventTransitions[S, E, C]{
		eventTransitions: make(map[E]*Transition[S, E, C]),
	}
}

// Put 添加一个 transition
func (et *EventTransitions[S, E, C]) Put(event E, transition *Transition[S, E, C]) error {
	et.mu.Lock()
	defer et.mu.Unlock()

	if _, ok := et.eventTransitions[event]; ok {
		return et.verify(et.eventTransitions[event], transition)
	}
	et.eventTransitions[event] = transition
	return nil
}

// Get 获取某个 event 的 transition
func (et *EventTransitions[S, E, C]) Get(event E) *Transition[S, E, C] {
	et.mu.RLock()
	defer et.mu.RUnlock()
	return et.eventTransitions[event]
}

// AllTransitions 获取所有 transition
func (et *EventTransitions[S, E, C]) AllTransitions() []*Transition[S, E, C] {
	et.mu.RLock()
	defer et.mu.RUnlock()
	var all []*Transition[S, E, C]
	for _, transition := range et.eventTransitions {
		all = append(all, transition)
	}
	return all
}

// verify event 只允许一个 transition
func (et *EventTransitions[S, E, C]) verify(existing *Transition[S, E, C], newT *Transition[S, E, C]) error {
	if existing.event == newT.event {
		return fmt.Errorf("transition %+v already exists, you cannot add another one", existing)
	}
	return nil
}
