package renderer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/thneutral/go-trpg-game/internal/character"
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
	TileSize     utils.Vector2D
	ScreenCenter utils.Vector2D
	ScreenSize   utils.Vector2D
	Field        *field.Field
}

func (renderer *PerpendicularRenderer) DrawField(screen *ebiten.Image, index utils.Vector2D) {
	offset := renderer.ScreenCenter.Subtract(renderer.ScreenSize.Scale(0.5))
	pos := offset.Add(index.Multiply(renderer.TileSize))
	tilePosition := pos.Scale(renderer.Zoom).Subtract(utils.Vector2D{X: 1, Y: 1})
	tileSize := renderer.TileSize.Scale(renderer.Zoom).Add(utils.Vector2D{X: 1, Y: 1})
	vector.DrawFilledRect(screen, float32(tilePosition.X), float32(tilePosition.Y), float32(tileSize.X), float32(tileSize.Y), renderer.Field.GetColor(index), false)
	if char := renderer.Field.GetCharacter(index); char != nil {
		switch char.Type {
		case character.MELEE:
			{
				circleCenter := pos.Add(renderer.TileSize.Scale(0.5)).Scale(renderer.Zoom)
				circleRadius := float32(renderer.TileSize.Scale(0.5).Scale(renderer.Zoom).GetSmallestCoordinate()) / 2.0
				var color color.Color
				if char.IsMine {
					color = field.GREEN_PURE
				} else {
					color = field.RED_PURE
				}
				vector.DrawFilledCircle(screen, float32(circleCenter.X), float32(circleCenter.Y), circleRadius, color, false)
			}
		}
	}
}

func (renderer *PerpendicularRenderer) DrawAllFields(screen *ebiten.Image) {
	for index, _ := range renderer.Field.Tiles {
		x, y := utils.OneDimToTwoDim(index, renderer.Field.XSize)
		pos := utils.Vector2D{X: x, Y: y}
		renderer.DrawField(screen, pos)
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

func GetRenderer(tileSize utils.Vector2D, field *field.Field, screenWidth, screenHeight int) *PerpendicularRenderer {
	return &PerpendicularRenderer{
		Zoom:     1,
		TileSize: tileSize,
		Field:    field,
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
