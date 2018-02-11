package labyrinth

import "strconv"

type FpsText struct {
	*Text
	time   float64
	update float64
}

func NewFpsText(x, y int, fg, bg Attr, update float64) *FpsText {
	return nil
}

func (f *FpsText) Draw(s *Screen) {

}

func cubeIndex(x int, points [5]int) int {
	return 0
}

func RgbTo256Color(r, g, b int) Attr {
	return 0
}