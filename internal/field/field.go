package field

import (
	"container/list"
	"image/color"

	"github.com/thneutral/go-trpg-game/internal/utils"
)

var (
	BLACK = color.Black
	WHITE = color.White

	RED   = color.RGBA{R: 255, A: 255}
	GREEN = color.RGBA{G: 255, A: 255}
	BLUE  = color.RGBA{B: 255, A: 255}

	THAT_BLUE = color.RGBA{R: 0x49, G: 0xE0, B: 0xE4, A: 0xFF}
)

func GetNewField(xsize, ysize int) *Field {
	arr := make([]Tile, 0)
	for y := range ysize {
		for x := range xsize {
			var t Tile
			t.Pos = utils.TwoDimToOneDim(x, y, xsize)
			t.InnerColor = WHITE
			t.BorderColor = BLACK
			arr = append(arr, t)
		}
	}

	return &Field{
		XSize: xsize,
		YSize: ysize,
		Tiles: arr,
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
	Tiles []Tile
}

func (field *Field) GetBorderColor(pos utils.Vector2D) color.Color {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	return field.Tiles[index].BorderColor
}

func (field *Field) GetInnerColor(pos utils.Vector2D) color.Color {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	return field.Tiles[index].InnerColor
}

func (field *Field) ChangeBorderColor(pos utils.Vector2D, color color.Color) {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	field.Tiles[index].BorderColor = color
}

func (field *Field) ChangeInnerColor(pos utils.Vector2D, color color.Color) {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	field.Tiles[index].InnerColor = color
}

func (field *Field) CreateSinglePath(origin utils.Vector2D, end utils.Vector2D, tilesToUse []utils.Vector2D) []utils.Vector2D {
	tileSet := make(map[utils.Vector2D]bool)
	visited := make(map[utils.Vector2D]bool)
	for _, tile := range tilesToUse {
		tileSet[tile] = true
		visited[tile] = false
	}
	if !tileSet[origin] || !tileSet[end] {
		return []utils.Vector2D{}
	}

	queue := list.New()
	queue.PushBack([]utils.Vector2D{origin})
	path := make([]utils.Vector2D, 0)
	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).([]utils.Vector2D)
		currentPosition := currentPath[len(currentPath)-1]

		if currentPosition.Equals(end) {
			path = currentPath
			break
		}

		for _, dir := range utils.GetAllDirections() {
			next := currentPosition.Add(dir)

			if !next.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
				continue
			}
			if visited[next] {
				continue
			}
			if !tileSet[next] {
				continue
			}
			nextPath := append([]utils.Vector2D{}, currentPath...)
			nextPath = append(nextPath, next)
			queue.PushBack(nextPath)
		}
	}
	if len(path) == 0 {
		return []utils.Vector2D{}
	}
	for _, pos := range path {
		field.ChangeInnerColor(pos, THAT_BLUE)
	}
	field.ChangeInnerColor(origin, RED)
	field.ChangeInnerColor(end, RED)
	return path
}

func (field *Field) CreateAllPaths(origin utils.Vector2D, len int) []utils.Vector2D {
	tiles := &[]utils.Vector2D{}
	field.computePathsInternal(origin, len, 0, THAT_BLUE, tiles)
	field.ChangeInnerColor(origin, RED)
	return *tiles
}

func (field *Field) DeleteAllPaths(origin utils.Vector2D, len int) {
	field.computePathsInternal(origin, len, 0, WHITE, nil)
	field.ChangeInnerColor(origin, WHITE)
}

func (field *Field) computePathsInternal(origin utils.Vector2D, len int, depth int, color color.Color, tiles *[]utils.Vector2D) {
	if depth > len {
		return
	}
	field.ChangeInnerColor(origin, color)
	for _, dirs := range utils.GetAllDirections() {
		new := origin.Add(dirs)
		if !new.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
			continue
		}
		if tiles != nil {
			utils.AddUnique(tiles, new)
		}
		field.computePathsInternal(new, len, depth+1, color, tiles)
	}
}
