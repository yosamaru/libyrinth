package labyrinth

import (
	"time"
	"github.com/nsf/termbox-go"
	"fmt"
)

// 一个游戏应该包含四部分
// 1、屏幕展示 2、输入操作 3、容错检测 4、日志记录
type Game struct {
	screen *Screen
	input  *input
	debug  bool
	logs   []string
}

// 创建新的游戏
func NewGame() *Game {
	g := Game{
		screen: NewScreen(),
		input:  newInput(),
		logs:   make([]string, 0),  // 创建一个指定元素与类型，长度的Slice
	}

	return &g
}

// 获取game的Screen
func (g *Game) Screen() *Screen {
	return g.screen
}

// 设置Screen
func (g *Game) SetScreen(s *Screen) {
	g.screen = s
	g.screen.resize(termbox.Size())
}

func (g *Game) DebugOn() bool {
	return g.debug
}

// 设置debug
func (g *Game) SetDebugOn(debugOn bool) {
	g.debug = debugOn
}

// log添加的log的信息
func (g *Game) Log(log string, items ...interface{}) {
	toLog := "[" + time.Now().Format(time.StampMilli) + "]" + fmt.Sprintf(log, items...)
	g.logs = append(g.logs, toLog)
}

// 打印log
func (g *Game) dumpLogs() {
	if g.debug {
		fmt.Println("===logs: ===")
		for _, l := range g.logs {
			fmt.Println(l)
		}
		fmt.Println("============")
	}
}

// 设置结束的快捷键
func (g *Game) SetEndKey(key Key) {
	g.input.endKey = termbox.Key(key)
}

func (g *Game) Start() {

}
