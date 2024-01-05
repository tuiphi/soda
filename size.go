package soda

import (
	"fmt"
	"math"
)

type Size struct {
	Width, Height int
}

func (s Size) String() string {
	return fmt.Sprint(s.Width, "x", s.Height)
}

func (s Size) SplitHorizontal(leftRatio float64) (Size, Size) {
	width := s.Width
	left := int(math.Round(float64(width) * leftRatio))
	right := width - left

	return Size{
			Width:  left,
			Height: s.Height,
		}, Size{
			Width:  right,
			Height: s.Height,
		}
}

func (s Size) SplitVertical(topRatio float64) (Size, Size) {
	left, right := s.swapped().SplitHorizontal(topRatio)

	return left.swapped(), right.swapped()
}

func (s Size) swapped() Size {
	return Size{
		Width:  s.Height,
		Height: s.Width,
	}
}
