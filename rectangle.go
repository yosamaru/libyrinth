package labyrinth

//创建矩形
type Rectangle struct {
	x      int
	y      int
	width  int
	height int
	color  Attr
}

func NewRectangle(x, y, w, h int, color Attr) *Rectangle {
	r := Rectangle{x: x, y: y, width: w, height: h, color: color}
	return &r
}

func (r *Rectangle) Draw(s *Screen) {
	for i := 0; i < r.width; i++ {
		for j := 0; j < r.height; j++ {
			s.RenderCell(r.x+i, r.y+j, &Cell{Bg: r.color, Ch: ' '})
		}
	}
}

func (r *Rectangle) Tick(ev Event) {}

func (r *Rectangle) Size() (int, int) {
	return r.width, r.height
}

func (r *Rectangle) Position() (int, int) {
	return r.x, r.y
}

func (r *Rectangle) SetPosition (x, y int) {
	r.x = x
	r.y = y
}

// 设置宽高
func (r *Rectangle) SetSize(w, h int) {
	r.width = w
	r.height = h
}

func (r *Rectangle) Color() Attr {
	return r.color
}

func (r *Rectangle) SetColor(color Attr) {
	r.color = color
}