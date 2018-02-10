package labyrinth

import termbox "github.com/nsf/termbox-go"

//幕布设计
type Canvas [][]Cell

func NewCanvas(width, height int) Canvas {

}

// 比较是否是同一个Canvas
func (canvas *Canvas) equals(oldCanvas Canvas) bool {

}

// 根据输入的信息创建新的Canvas
func CanvasFromString(str string) Canvas {

}

//创建绘制的接口
type Drawable interface {
}

// 创建物理接口
type Physical interface {
}

// 创建动态的physical
type DynamicPhysical interface {
}

//比较最小值
func min(a, b int) int {

}

// 比较最大值
func max(a, b int) int {

}

type Cell struct {
	Fg Attr
	Bg Attr
	Ch rune
}

func (c *Cell) equals(c2 *Cell) bool {

}

type Event struct {
	Type   EventType
	Key    Key
	Ch     rune
	Mod    Modifier
	Err    error
	MouseX int
	MouseY int
}

func convertEvent(ev termbox.Event) Event {

}

type (
	Attr      uint16
	Key       uint16
	Modifier  uint8
	EventType uint8
)
