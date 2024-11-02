package field

import (
	"image/color"

	"github.com/thneutral/go-trpg-game/internal/utils"
)

func GetNewField(xsize, ysize int) *Field {
	// colors := []color.RGBA{color.RGBA{R: 255, A: 255}, color.RGBA{G: 255, A: 255}, color.RGBA{B: 255, A: 255}}
	arr := make([]Tile, 0)
	for y := range ysize {
		for x := range xsize {
			var t Tile
			t.Pos = utils.TwoDimToOneDim(x, y, xsize)
			t.InnerColor = color.White
			t.BorderColor = color.Black
			arr = append(arr, t)
		}
	}

	return &Field{
		XSize: xsize,
		YSize: ysize,
		Arr:   arr,
	}
}

type Tile struct {
	Pos         int
	InnerColor  color.Color
	BorderColor color.Color
}

type Field struct {
	XSize int
	YSize int
	Arr   []Tile
}

func (f *Field) ChangeColor(x, y int, color color.Color) {
	index := utils.TwoDimToOneDim(x, y, f.XSize)
	f.Arr[index].InnerColor = color
}
