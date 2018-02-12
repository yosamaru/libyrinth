package labyrinth

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
  level := BaseLevel{Entities: make([]Drawable, 0), bg: bg}
  return &level
}

// 标号
func (l *BaseLevel) Tick(ev Event) {
  if ev.Type != EventNone {
    for _, e := range l.Entities {
      e.Tick(ev)
    }
  }

  colls := make([]Physical, 0)
  dynamics := make([]DynamicPhysical, 0)
  for _, e := range l.Entities {
    // TODO 将值e转为空接口类型,并用comma-ok断言是不是指定的Physical类型
    if p, ok := interface{}(e).(Physical); ok {
      colls = append(colls, p)
    }
    if p, ok := interface{}(e).(DynamicPhysical); ok {
      dynamics = append(dynamics, p)
    }
  }
  jobs := make(chan DynamicPhysical, len(dynamics))
  results := make(chan int, len(dynamics))
  for w := 0; w <= len(dynamics)/3; w++ {
    go checkCollisionsWorker(colls, jobs, results)
  }
  for _, p := range dynamics {
    jobs <- p
  }
  close(jobs)
  for r := 0; r < len(dynamics); r++ {
    <-results
  }
}

// 绘制背景
func (l *BaseLevel) DrawBackground(s *Screen) {
  for i, row := range s.canvas {
    for j, _ := range row  {
      s.canvas[i][j] = l.bg
    }
  }
}

// 画笔
func (l *BaseLevel) Draw(s *Screen) {
  offx, offy := s.offset()
  s.setOffset(l.offsetx, l.offsety)
  for _, e := range l.Entities {
    e.Draw(s)
  }
  s.setOffset(offx, offy)
}

func (l *BaseLevel) AddEntity(d Drawable) {
  l.Entities = append(l.Entities, d)
}

func (l *BaseLevel) RemoveEntity(d Drawable) {
  for i, elem := range l.Entities {
    if elem == d {
      l.Entities = append(l.Entities[:i], l.Entities[i+1:]...)
      return
    }
  }
}

// 位移
func (l *BaseLevel) Offset() (int, int) {
  return l.offsetx, l.offsety
}

func (l *BaseLevel) SetOffset(x, y int) {
  l.offsetx, l.offsety = x, y
}

func checkCollisionsWorker(ps []Physical, jobs <-chan DynamicPhysical, results chan<- int) {
  for p := range jobs {
    for _, c := range ps {
      if c  == p {
        continue
      }
      px, py := p.Position()
      cx, cy := c.Position()
      pw, ph := p.Size()
      cw, ch := c.Size()
      if px < cx+cw && px+pw > cx &&
        py < cy+ch && py+ph > cy {
          p.Collide(c)
      }
    }
    results <- 1
  }
}





