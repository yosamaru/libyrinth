package labyrinth

import "database/sql/driver"

// 等级操作
type Level interface {
  DrawBackground(*Screen)
  AddEntity(Drawable)
  RemoveEntity(Drawable)
  Draw(*Screen)
  Tick(Event)
}

// 基础的等级设置
type BaseLevel struct {
  Entities []Drawable
  bg       Cell
  offsetx  int
  offsety  int
}

// 创建新的基础等级
func NewBaseLevel(bg Cell) *BaseLevel {
  return nil
}

// 标号
func (l *BaseLevel) Tick(ev Event) {

}

// 绘制背景
func (l *BaseLevel) DrawBackground(s *Screen) {

}

// 画笔
func (l *BaseLevel) Draw(s *Screen) {

}

func (l *BaseLevel) AddEntity(d Drawable) {

}

func (l *BaseLevel) RemoveEntity(d Drawable) {

}

// 位移
func (l *BaseLevel) Offset() (int, int) {
  return 0, 0
}

func (l *BaseLevel) SetOffset(x, y int) {

}

func checkCollisionsWorker(ps []Physical, jobs <-chan DynamicPhysical, results chan<- int) {

}





