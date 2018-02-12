package labyrinth

type Entity struct {
	canvas Canvas
	x      int
	y      int
	width  int
	height int
}

func NewEntity(x, y, width, height int) *Entity {
	canvas := NewCanvas(width, height)
	e := Entity{x: x, y: y, width: width, height: height, canvas: canvas}
	return &e
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
	for i := 0; i < e.width; i++ {
		for j := 0; j < e.height; j++ {
			s.RenderCell(e.x+i, e.y+j, &e.canvas[i][j])
		}
	}
}

func (e *Entity) Tick(ev Event) {}

func (e *Entity) Position() (int, int) {
	return e.x, e.y
}

func (e *Entity) Size() (int, int) {
	return e.width, e.height
}

func (e *Entity) SetPosition(x, y int) {
	e.x = x
	e.y = y
}

func (e *Entity) SetCell(x, y int, c *Cell) {
	renderCell(&e.canvas[x][y], c)
}

// Cell的填充
func (e *Entity) Fill(c *Cell)  {
	for  i := range e.canvas {
		for j := range e.canvas[i] {
			renderCell(&e.canvas[i][j], c)
		}
	}
}

func (e *Entity) ApplyCanvas(c *Canvas) {
	for i := 0; i < min(len(e.canvas), len(*c)); i++ {
		for j := 0; j < min(len(e.canvas[0]), len((*c)[0])); j++ {
			renderCell(&e.canvas[i][j], &(*c)[i][j])
		}
	}
}

func (e *Entity) SetCanvas(c *Canvas) {
	e.width = len(*c)
	e.height = len((*c)[0])
	e.canvas = *c
}