package draw

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func SquareFill(screen *ebiten.Image, x1, y1, x2, y2 int, clr color.RGBA) {
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			screen.Set(x, y, clr)
		}
	}
}
