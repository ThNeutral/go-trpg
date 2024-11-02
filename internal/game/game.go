package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/thneutral/go-trpg-game/internal/field"
	"github.com/thneutral/go-trpg-game/internal/renderer"
	"github.com/thneutral/go-trpg-game/internal/utils"
)

const (
	ZOOM_RATE = 1.2

	MAX_ZOOM = ZOOM_RATE * ZOOM_RATE * ZOOM_RATE
	MIN_ZOOM = 1 / ZOOM_RATE / ZOOM_RATE / ZOOM_RATE
)

func GetNewGame(field *field.Field, renderer *renderer.PerpendicularRenderer) *Game {
	x, y := ebiten.CursorPosition()
	return &Game{
		Field:    field,
		Renderer: renderer,
		OldCursorPosition: utils.Vector2D{
			X: float32(x),
			Y: float32(y),
		},
	}
}

type Game struct {
	Field             *field.Field
	Renderer          *renderer.PerpendicularRenderer
	OldCursorPosition utils.Vector2D
}

func (g *Game) Update() error {
	xCursor, yCursor := ebiten.CursorPosition()
	_, yWheel := ebiten.Wheel()
	if yWheel > 0 && g.Renderer.Zoom < MAX_ZOOM {
		g.Renderer.Zoom *= ZOOM_RATE
	} else if yWheel < 0 && g.Renderer.Zoom > MIN_ZOOM {
		g.Renderer.Zoom /= ZOOM_RATE
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton1) {
		g.Renderer.ScreenCenter.X += (float32(xCursor) - g.OldCursorPosition.X) / g.Renderer.Zoom
		g.Renderer.ScreenCenter.Y += (float32(yCursor) - g.OldCursorPosition.Y) / g.Renderer.Zoom
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		xIndex, yIndex := g.Renderer.ScreenToFieldIndex(xCursor, yCursor)
		if xIndex != -1 && yIndex != -1 {
			g.Field.ChangeColor(xIndex, yIndex, color.RGBA{R: 0x49, G: 0xE0, B: 0xE4, A: 0xFF})
		}
	}
	g.OldCursorPosition = utils.Vector2D{
		X: float32(xCursor),
		Y: float32(yCursor),
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// screen.Fill(co]lor.RGBA{R: 255, A: 255})
	screen.Clear()
	g.Renderer.DrawField(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
