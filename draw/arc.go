package draw

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func ArcFill(screen, draw *ebiten.Image, x, y, radius int, clr color.RGBA) {
	x32 := float32(x)
	y32 := float32(y)
	r32 := float32(radius)

	r := float32(clr.R) / 255
	g := float32(clr.G) / 255
	b := float32(clr.B) / 255
	a := float32(clr.A) / 255

	var path vector.Path
	path.MoveTo(x32, y32)
	path.Arc(x32, y32, r32, 0, math.Pi*2, vector.Clockwise)

	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		v := &vs[i]
		v.ColorR = r
		v.ColorG = g
		v.ColorB = b
		v.ColorA = a
	}

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}

	screen.DrawTriangles(vs, is, draw, op)
}
