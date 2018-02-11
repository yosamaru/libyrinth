package labyrinth

import (
	termbox "github.com/nsf/termbox-go"
	"strings"
)

//幕布设计
type Canvas [][]Cell

func NewCanvas(width, height int) Canvas {
	return nil
}

// 比较是否是同一个Canvas
func (canvas *Canvas) equals(oldCanvas Canvas) bool {
	return false
}

// 根据输入的信息创建新的Canvas
func CanvasFromString(str string) Canvas {
	return nil
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
	return 0
}

// 比较最大值
func max(a, b int) int {
	return 0
}
// 创建格子类型
type Cell struct {
	Fg Attr
	Bg Attr
	Ch rune
}

// 设置每一个格子
func (c *Cell) equals(c2 *Cell) bool {
	return false
}

// 创建事件实体
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
	return Event{}
}

// 创建多个不同的类型
type (
	Attr      uint16
	Key       uint16
	Modifier  uint8
	EventType uint8
)

const (
	EventKey EventType = iota
	EventResize
	EventMouse
	EventError
	EventInterrupt
	EventRaw
	EventNone
)

const (
	ColorDefault Attr = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	ArrrBold Attr = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

const ModAltModifier  = 0x01

const (
	KeyF1 Key = 0xFFFF - iota
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	key_min
	MouseLeft
	MouseMiddle
	MouseRight
	MouseRelease
	MouseWheelUp
	MouseWheelDown
)

const (
	KeyCtrlTilde      Key = 0x00
	KeyCtrl2          Key = 0x00
	KeyCtrlSpace      Key = 0x00
	KeyCtrlA          Key = 0x01
	KeyCtrlB          Key = 0x02
	KeyCtrlC          Key = 0x03
	KeyCtrlD          Key = 0x04
	KeyCtrlE          Key = 0x05
	KeyCtrlF          Key = 0x06
	KeyCtrlG          Key = 0x07
	KeyBackspace      Key = 0x08
	KeyCtrlH          Key = 0x08
	KeyTab            Key = 0x09
	KeyCtrlI          Key = 0x09
	KeyCtrlJ          Key = 0x0A
	KeyCtrlK          Key = 0x0B
	KeyCtrlL          Key = 0x0C
	KeyEnter          Key = 0x0D
	KeyCtrlM          Key = 0x0D
	KeyCtrlN          Key = 0x0E
	KeyCtrlO          Key = 0x0F
	KeyCtrlP          Key = 0x10
	KeyCtrlQ          Key = 0x11
	KeyCtrlR          Key = 0x12
	KeyCtrlS          Key = 0x13
	KeyCtrlT          Key = 0x14
	KeyCtrlU          Key = 0x15
	KeyCtrlV          Key = 0x16
	KeyCtrlW          Key = 0x17
	KeyCtrlX          Key = 0x18
	KeyCtrlY          Key = 0x19
	KeyCtrlZ          Key = 0x1A
	KeyEsc            Key = 0x1B // No longer supported
	KeyCtrlLsqBracket Key = 0x1B
	KeyCtrl3          Key = 0x1B
	KeyCtrl4          Key = 0x1C
	KeyCtrlBackslash  Key = 0x1C
	KeyCtrl5          Key = 0x1D
	KeyCtrlRsqBracket Key = 0x1D
	KeyCtrl6          Key = 0x1E
	KeyCtrl7          Key = 0x1F
	KeyCtrlSlash      Key = 0x1F
	KeyCtrlUnderscore Key = 0x1F
	KeySpace          Key = 0x20
	KeyBackspace2     Key = 0x7F
	KeyCtrl8          Key = 0x7F
)