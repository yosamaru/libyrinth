package labyrinth

import "github.com/nsf/termbox-go"

type input struct {
	endKey termbox.Key
	eventQ chan termbox.Event
	ctrl   chan bool
}

// 创建新的输入
func newInput() *input {
	i := input{eventQ: make(chan termbox.Event),
		ctrl: make(chan bool, 2),
		endKey: termbox.KeyCtrlC}
	return &i
}

// 开始输入
func (i *input) start() {
	go poll(i)
}

// 停止输入
func (i *input) stop() {
	i.ctrl <- true
}

// 检测
func poll(i *input) {
loop:
	for {
		select {
		case <-i.ctrl:
			break loop
		default:
			i.eventQ <- termbox.PollEvent()
		}
	}
}
