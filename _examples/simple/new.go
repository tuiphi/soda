package main

func New(n int) *State {
	return &State{
		n:      n,
		keyMap: NewKeyMap(),
	}
}
