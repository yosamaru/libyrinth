package labyrinth

type Text struct {
	x      int
	y      int
	fg     Attr
	bg     Attr
	text   []rune
	canvas []Cell
}

func NewText(x, y int, text string, fg, bg Attr) {

}

func (t *Text) Tick(ev Event) {}

func (t *Text) Draw(s *Screen) {

}

func (t *Text) Position() (int, int) {
	return 0, 0
}

func (t *Text) Size() (int, int) {
	return 0, 0
}

func (t *Text) SetPosition(x, y int) {

}

func (t *Text) Text() string {
	return ""
}

func (t *Text) SetText(text string) {

}

func (t *Text) Color() (Attr, Attr) {
	return 0, 0
}

func (t *Text) SetColor(fg, bg Attr) {

}