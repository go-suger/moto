package moto

func getState[S, E comparable, C any](stateMap map[S]*State[S, E, C], stateVal S) *State[S, E, C] {
	state, ok := stateMap[stateVal]
	if !ok {
		state = newState[S, E, C](stateVal)
		stateMap[stateVal] = state
	}

	return state
}
