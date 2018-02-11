package labyrinth

import "github.com/nsf/termbox-go"

type input struct {
  endKey termbox.key
  eventQ chan termbox.Event
  ctrl   chan bool
}

// 创建新的输入
func newInput() *input {
  return nil
}

// 开始输入
func (i *input) start() {

}

// 停止输入
func (i *input) stop() {

}

// 检测
func poll(i *input) {

}
