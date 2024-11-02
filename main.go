package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/thneutral/go-trpg-game/internal/field"
	"github.com/thneutral/go-trpg-game/internal/game"
	"github.com/thneutral/go-trpg-game/internal/renderer"
	"github.com/thneutral/go-trpg-game/internal/utils"
)

const (
	X_SIZE = 100
	Y_SIZE = 100

	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720

	WINDOW_NAME = "trpg"
)

var (
	TILE_SIZE   = utils.Vector2D{X: 22, Y: 22}
	BORDER_SIZE = utils.Vector2D{X: 3, Y: 3}
)

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle(WINDOW_NAME)

	field := field.GetNewField(X_SIZE, Y_SIZE)
	renderer := renderer.GetRenderer(TILE_SIZE, BORDER_SIZE, field, SCREEN_WIDTH, SCREEN_HEIGHT)

	game := game.GetNewGame(field, renderer)
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
