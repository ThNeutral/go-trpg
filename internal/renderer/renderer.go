package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/thneutral/go-trpg-game/internal/field"
	"github.com/thneutral/go-trpg-game/internal/utils"
)

const (
	PERPENDICULAR_RENDERER = iota
)

type IRenderer interface {
	DrawField(screen *ebiten.Image)
}

type PerpendicularRenderer struct {
	Zoom         float32
	BorderSize   utils.Vector2D
	TileSize     utils.Vector2D
	ScreenCenter utils.Vector2D
	ScreenSize   utils.Vector2D
	Field        *field.Field
}

func (renderer *PerpendicularRenderer) DrawField(screen *ebiten.Image, index utils.Vector2D) {
	offset := renderer.ScreenCenter.Subtract(renderer.ScreenSize.Scale(0.5))
	pos := offset.Add(index.Multiply(renderer.TileSize))
	borderPos := pos.Scale(renderer.Zoom)
	borderSize := renderer.TileSize.Add(renderer.BorderSize).Scale(renderer.Zoom)
	innerPos := pos.Add(renderer.BorderSize.Scale(0.5)).Scale(renderer.Zoom)
	innerSize := renderer.TileSize.Scale(renderer.Zoom)
	vector.DrawFilledRect(screen, float32(borderPos.X), float32(borderPos.Y), float32(borderSize.X), float32(borderSize.Y), renderer.Field.GetBorderColor(index), false)
	vector.DrawFilledRect(screen, float32(innerPos.X), float32(innerPos.Y), float32(innerSize.X), float32(innerSize.Y), renderer.Field.GetInnerColor(index), false)

}

func (renderer *PerpendicularRenderer) DrawAllFields(screen *ebiten.Image) {
	offset := renderer.ScreenCenter.Subtract(renderer.ScreenSize.Scale(0.5))
	for _, tile := range renderer.Field.Tiles {
		x, y := utils.OneDimToTwoDim(tile.Pos, renderer.Field.XSize)
		index := utils.Vector2D{X: x, Y: y}
		pos := offset.Add(index.Multiply(renderer.TileSize))
		borderPos := pos.Scale(renderer.Zoom)
		borderSize := renderer.TileSize.Add(renderer.BorderSize).Scale(renderer.Zoom)
		innerPos := pos.Add(renderer.BorderSize.Scale(0.5)).Scale(renderer.Zoom)
		innerSize := renderer.TileSize.Scale(renderer.Zoom)
		vector.DrawFilledRect(screen, float32(borderPos.X), float32(borderPos.Y), float32(borderSize.X), float32(borderSize.Y), renderer.Field.GetBorderColor(index), false)
		vector.DrawFilledRect(screen, float32(innerPos.X), float32(innerPos.Y), float32(innerSize.X), float32(innerSize.Y), renderer.Field.GetInnerColor(index), false)
	}
}

func (renderer *PerpendicularRenderer) ScreenToFieldIndex(screen utils.Vector2D) utils.Vector2D {
	offset := renderer.ScreenCenter.Subtract(renderer.ScreenSize.Scale(0.5))
	distance := screen.Scale(1 / renderer.Zoom).Subtract(offset)
	index := distance.Divide(renderer.TileSize)
	if index.X < 0 || index.X >= renderer.Field.XSize {
		index.X = -1
	}
	if index.Y < 0 || index.Y >= renderer.Field.YSize {
		index.Y = -1
	}
	return index
}

// Doesn't work
// func (renderer *PerpendicularRenderer) selectTilesToDraw() []field.Tile {
// 	tiles := make([]field.Tile, 0)
// 	start := utils.Vector2D{
// 		X: renderer.ScreenCenter.X - renderer.ScreenSize.X/2/renderer.Zoom,
// 		Y: renderer.ScreenCenter.Y - renderer.ScreenSize.Y/2/renderer.Zoom,
// 	}
// 	end := utils.Vector2D{
// 		X: renderer.ScreenCenter.X + renderer.ScreenSize.X/2/renderer.Zoom,
// 		Y: renderer.ScreenCenter.Y + renderer.ScreenSize.Y/2/renderer.Zoom,
// 	}
// 	for _, tile := range renderer.Field.Arr {
// 		x, y := utils.OneDimToTwoDim(tile.Pos, renderer.Field.XSize)
// 		offset := utils.Vector2D{
// 			X: renderer.ScreenCenter.X - renderer.ScreenSize.X/2/renderer.Zoom,
// 			Y: renderer.ScreenCenter.Y - renderer.ScreenSize.Y/2/renderer.Zoom,
// 		}
// 		pos := utils.Vector2D{
// 			X: offset.X + float32(x)*renderer.TileSize.X,
// 			Y: offset.Y + float32(y)*renderer.TileSize.Y,
// 		}
// 		if pos.X < start.X || pos.X > end.X || pos.Y < start.Y || pos.Y > end.Y {
// 			continue
// 		}
// 		tiles = append(tiles, tile)
// 	}
// 	return tiles
// }

func GetRenderer(tileSize utils.Vector2D, borderSize utils.Vector2D, field *field.Field, screenWidth, screenHeight int) *PerpendicularRenderer {
	return &PerpendicularRenderer{
		Zoom:       1,
		BorderSize: borderSize,
		TileSize:   tileSize,
		Field:      field,
		ScreenCenter: utils.Vector2D{
			X: screenWidth / 2,
			Y: screenHeight / 2,
		},
		ScreenSize: utils.Vector2D{
			X: screenWidth,
			Y: screenHeight,
		},
	}
}
