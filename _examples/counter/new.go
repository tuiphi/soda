package main

func New() *State {
	return &State{
		counter: 0,
		keyMap:  DefaultKeyMap(),
	}
}
