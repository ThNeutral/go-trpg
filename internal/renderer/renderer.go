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

func (renderer *PerpendicularRenderer) DrawField(screen *ebiten.Image) {
	offset := utils.Vector2D{
		X: renderer.ScreenCenter.X - renderer.ScreenSize.X/2,
		Y: renderer.ScreenCenter.Y - renderer.ScreenSize.Y/2,
	}
	for _, tile := range renderer.Field.Arr {
		x, y := utils.OneDimToTwoDim(tile.Pos, renderer.Field.XSize)
		pos := utils.Vector2D{
			X: offset.X + float32(x)*renderer.TileSize.X,
			Y: offset.Y + float32(y)*renderer.TileSize.Y,
		}
		vector.DrawFilledRect(screen, pos.X*renderer.Zoom, pos.Y*renderer.Zoom, (renderer.TileSize.X+renderer.BorderSize.X)*renderer.Zoom, (renderer.TileSize.Y+renderer.BorderSize.Y)*renderer.Zoom, tile.BorderColor, false)
		vector.DrawFilledRect(screen, (pos.X+renderer.BorderSize.X/2)*renderer.Zoom, (pos.Y+renderer.BorderSize.Y/2)*renderer.Zoom, renderer.TileSize.X*renderer.Zoom, renderer.TileSize.Y*renderer.Zoom, tile.InnerColor, false)
	}
}

func (renderer *PerpendicularRenderer) ScreenToFieldIndex(xScreen, yScreen int) (int, int) {
	offset := utils.Vector2D{
		X: renderer.ScreenCenter.X - renderer.ScreenSize.X/(2),
		Y: renderer.ScreenCenter.Y - renderer.ScreenSize.Y/(2),
	}
	xDistance := (float32(xScreen) / renderer.Zoom) - offset.X
	yDistance := (float32(yScreen) / renderer.Zoom) - offset.Y
	xIndex := xDistance / renderer.TileSize.X
	yIndex := yDistance / renderer.TileSize.Y
	if xIndex < 0 || xIndex >= float32(renderer.Field.XSize) {
		xIndex = -1
	}
	if yIndex < 0 || yIndex >= float32(renderer.Field.YSize) {
		yIndex = -1
	}

	return int(xIndex), int(yIndex)
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
			X: float32(screenWidth / 2),
			Y: float32(screenHeight / 2),
		},
		ScreenSize: utils.Vector2D{
			X: float32(screenWidth),
			Y: float32(screenHeight),
		},
	}
}
