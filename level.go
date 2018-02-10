package labyrinth

// 等级操作
type Level interface {
  DrawBackground(*Screen)
  AddEntity(Drawable)
  RemoveEntity(Drawable)
  Draw(*Screen)
  Tick(Event)
}
