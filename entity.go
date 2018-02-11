package labyrinth

type Entity struct {
	canvas Canvas
	x      int
	y      int
	width  int
	height int
}

func NewEntity(x, y, width, height int) *Entity {
	return nil
}

func NewEntityFromCanvas(x, y int, c Canvas) *Entity {
	e := Entity{
		x:      x,
		y:      y,
		canvas: c,
		width:  len(c),
		height: len(c[0]),
	}
	return &e
}

func (e *Entity) Draw(s *Screen) {

}

func (e *Entity) Tick(ev Event) {}

func (e *Entity) Position() (int, int) {
	return 0, 0
}

func (e *Entity) Size() (int, int) {
	return e.width, e.height
}

func (e *Entity) SetPosition(x, y int) {
	e.x = x
	e.y = y
}

func (e *Entity) SetCell(x, y int, c *Cell) {

}

// Cell的填充
func (e *Entity) Fill(c *Cell)  {

}

func (e *Entity) ApplyCanvas(c *Canvas) {

}

func (e *Entity) SetCanvas(c *Canvas) {

}