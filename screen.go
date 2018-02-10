package labyrinth

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

}
