package labyrinth

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func colorGridFromFile(filename string) *[][]Attr {
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y
	colors := make([][]Attr, w)
	for i := range  colors {
		colors[i] = make([]Attr, h)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			if a < 1 {
				continue
			}
			R := int(r >> 8)
			G := int(g >> 8)
			B := int(b >> 8)
			colors[x-bounds.Min.X][y-bounds.Min.Y] = RgbTo256Color(R, G, B)
		}
	}
	return &colors
}

func BackgroundCanvasFromFile(filename string) *Canvas {
	colors := colorGridFromFile(filename)
	c := make(Canvas, len(*colors))
	for i := range c {
		c[i] = make([]Cell, len((*colors)[i]))
		for j := range c {
			c[i][j] = Cell{Bg: (*colors)[i][j]}
		}
	}
	return &c
}

func ForegroundCanvasFromFile(filename string) *Canvas {
	colors := colorGridFromFile(filename)
	c := make(Canvas, len(*colors))
	for i := range c {
		c[i] = make([]Cell, len((*colors)[i]))
		for j := range c[i] {
			c[i][j] = Cell{Fg: (*colors)[i][j]}
		}
	}
	return &c
}
