package soda

import "fmt"

type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprint(s.Width, "x", s.Height)
}
