package game

import (
	"slices"

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

const (
	GAME_STATE_PLAYING = iota
)

const (
	SELECTION_STATE_NONE = iota
	SELECTION_STATE_SELECTED
	SELECTION_STATE_PATH
)

func GetNewGame(field *field.Field, renderer *renderer.PerpendicularRenderer) *Game {
	x, y := ebiten.CursorPosition()
	return &Game{
		Field:    field,
		Renderer: renderer,
		OldCursorPosition: utils.Vector2D{
			X: x,
			Y: y,
		},
		GameState:      GAME_STATE_PLAYING,
		SelectionState: SELECTION_STATE_NONE,
	}
}

type Game struct {
	Field             *field.Field
	Renderer          *renderer.PerpendicularRenderer
	OldCursorPosition utils.Vector2D
	SelectionCenter   utils.Vector2D
	SelectedTiles     []utils.Vector2D
	GameState         int
	SelectionState    int
}

func (game *Game) Update() error {
	// Common data before
	xCursor, yCursor := ebiten.CursorPosition()
	_, yWheel := ebiten.Wheel()

	// Zoom
	if yWheel > 0 && game.Renderer.Zoom < MAX_ZOOM {
		game.Renderer.Zoom *= ZOOM_RATE
	} else if yWheel < 0 && game.Renderer.Zoom > MIN_ZOOM {
		game.Renderer.Zoom /= ZOOM_RATE
	}

	// Drag
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton1) {
		game.Renderer.ScreenCenter.X += int(float32(xCursor-game.OldCursorPosition.X) / game.Renderer.Zoom)
		game.Renderer.ScreenCenter.Y += int(float32(yCursor-game.OldCursorPosition.Y) / game.Renderer.Zoom)
	}

	// Selection
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		selectedTile := game.Renderer.ScreenToFieldIndex(utils.Vector2D{X: xCursor, Y: yCursor})
		if !selectedTile.IsAnyNegative() {
			switch game.SelectionState {
			case SELECTION_STATE_NONE:
				{
					game.SelectionState = SELECTION_STATE_SELECTED
					game.SelectedTiles = game.Field.CreateAllPaths(selectedTile, 5)
					game.SelectionCenter = selectedTile
				}
			case SELECTION_STATE_SELECTED:
				{
					game.Field.DeleteAllPaths(game.SelectionCenter, 5)
					if selectedTile.Equals(game.SelectionCenter) {
						game.SelectionState = SELECTION_STATE_NONE
					} else if slices.Contains(game.SelectedTiles, selectedTile) {
						game.SelectionState = SELECTION_STATE_PATH
						game.Field.CreateSinglePath(game.SelectionCenter, selectedTile, game.SelectedTiles)
					} else {
						game.SelectionState = SELECTION_STATE_NONE
					}
				}
			case SELECTION_STATE_PATH:
				{
					game.SelectionState = SELECTION_STATE_NONE
					game.Field.DeleteAllPaths(game.SelectionCenter, 5)
				}
			}
		}
	}

	// Common data after
	game.OldCursorPosition = utils.Vector2D{
		X: xCursor,
		Y: yCursor,
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	game.Renderer.DrawAllFields(screen)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
