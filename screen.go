package labyrinth

import "github.com/nsf/termbox-go"

// 幕布 等级以及screen的基础信息
type Screen struct {
	oldCanvas Canvas
	canvas    Canvas
	level     Level
	Entities  []Drawable
	width     int
	height    int
	delta     float64
	fps       float64
	offsetx   int
	offsety   int
	pixelMode bool
}

//创建新的Screen
func NewScreen() *Screen {
	s := Screen{Entities: make([]Drawable, 0)}
	s.canvas = NewCanvas(10, 10)
	return &s
}

func (s *Screen) Tick(ev Event) {
	if s.level != nil {
		s.level.Tick(ev)
	}

	if ev.Type != EventNone {
		for _, e := range s.Entities {
			e.Tick(ev)
		}
	}
}

func (s *Screen) Draw() {
	// 更新canvas
	s.canvas = NewCanvas(s.width, s.height)
	if s.level != nil {
		s.level.DrawBackground(s)
		s.level.Draw(s)
	}

	for _, e := range s.Entities {
		e.Draw(s)
	}

	if !s.canvas.equals(&s.oldCanvas) {
		if s.pixelMode {
			termboxPixel(&s.canvas)
		} else {
			termboxNormal(&s.canvas)
		}
		termbox.Flush()
	}
	s.oldCanvas = s.canvas
}

func (s *Screen) resize(w, h int) {
	s.width = w
	s.height = h
	if s.pixelMode {
		s.height *= 2
	}

	c := NewCanvas(s.width, s.height)

	for i := 0; i < min(s.width, len(s.canvas)); i ++ {
		for j := 0; j < min(s.height, len(s.canvas[0])); j++ {
			c[i][j] = s.canvas[i][j]
		}
	}
	s.canvas = c
}

func (s *Screen) Size() (int, int) {
	return s.width, s.height
}

func (s *Screen) SetLevel(l Level) {
	s.level = l
}

func (s *Screen) Level() Level {
	return s.level
}

func (s *Screen) AddEntity(d Drawable) {
	s.Entities = append(s.Entities, d)
}

func (s *Screen) RemoveEntity(d Drawable) {
	for i, elem := range s.Entities {
		if elem == d {
			s.Entities = append(s.Entities[:i], s.Entities[i+1:]...)
			return
		}
	}
}

func (s *Screen) TimeDelta() float64 {
	return s.delta
}

// 设置帧数
func (s *Screen) SetFps(f float64) {
	s.fps = f
}

func (s *Screen) RenderCell(x, y int, c *Cell) {
	newx := x + s.offsetx
	newy := y + s.offsety

	if newx >= 0 && newx < len(s.canvas) &&
		newy >= 0 && newy < len(s.canvas[0]) {
		renderCell(&s.canvas[newx][newy], c)
	}
}

func (s *Screen) EnablePixelModel() {
	s.pixelMode = true
}

func (s *Screen) offset() (int, int) {
	return s.offsetx, s.offsety
}

func (s *Screen) setOffset(x, y int) {
	s.offsetx, s.offsety = x, y
}

// 渲染Cell
func renderCell(old, new_ *Cell) {
	if new_.Ch != 0 {
		old.Ch = new_.Ch
	}

	if new_.Bg != 0 {
		old.Bg = new_.Bg
	}

	if new_.Fg != 0 {
		old.Fg = new_.Fg
	}
}

func termboxPixel(canvas *Canvas) {
	for i, col := range *canvas {
		for j := 0; j < len(col); j += 2 {
			cellBack := col[j]
			cellFront := col[j+1]
			termj := j / 2
			char := '\u2584'
			if cellFront.Bg == 0 {
				char = 0
			}
			termbox.SetCell(i, termj, char, termbox.Attribute(cellFront.Bg),
				termbox.Attribute(cellBack.Bg))
		}
	}
}

func termboxNormal(canvas *Canvas) {
	for i, col := range *canvas {
		for j, cell := range col {
			termbox.SetCell(i, j, cell.Ch,
				termbox.Attribute(cell.Fg),
				termbox.Attribute(cell.Bg))
		}
	}
}
