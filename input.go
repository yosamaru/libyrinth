package labyrinth

import "github.com/github.com/nsf/termbox-go"

type input struct {
  endKey termbox.key
  eventQ chan termbox.Event
  ctrl   chan bool
}
