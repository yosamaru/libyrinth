package labyrinth

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

}

// 获取game的Screen
func (g *Game) Screen() *Screen {

}

// 设置Screen
func (g *Game) SetScreen(s *Screen) {

}

func (g *Game) DebugOn() bool {

}

// 设置debug
func (g *Game) SetDebugOn(debugOn bool) {

}

// log设置
func (g *Game) Log(log string, items ...interface{}) {

}

// 无用的log
func (g *Game) dumpLogs() {

}

func (g *Game) SetEndKey(key Key) {

}

func (g *Game) Start() {

}
