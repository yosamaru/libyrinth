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
	return nil
}

func (r *Rectangle) Draw(s *Screen) {

}

func (r *Rectangle) Tick(ev Event) {}

func (r *Rectangle) Size() (int, int) {
	return 0, 0
}

func (r *Rectangle) Position() (int, int) {
	return 0, 0
}

func (r *Rectangle) SetPosition (x, y int) {

}

// 设置宽高
func (r *Rectangle) SetSize(w, h int) {

}

func (r *Rectangle) Color() Attr {
	return r.color
}

func (r *Rectangle) SetColor(color Attr) {

}