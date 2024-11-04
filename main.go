package main

import (
	"flag"
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

	MIN_POCKET_SIZE = 100
	SPAWN_SIZE      = 25

	WINDOW_NAME = "trpg"
)

var (
	TILE_SIZE = utils.Vector2D{X: 30, Y: 30}
)

func main() {
	genTypeFlag := flag.String("generation", "OPEN_SIMPLEX", "Select method to generate map.\n\tOPEN_SIMPLEX, PURE_RANDOM, PERLIN_NOISE")
	flag.Parse()

	ebiten.SetWindowSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	ebiten.SetWindowTitle(WINDOW_NAME)

	var genType int
	switch *genTypeFlag {
	case "OPEN_SIMPLEX":
		{
			genType = field.OPEN_SIMPLEX_NOISE_GENERATION
		}
	case "PURE_RANDOM":
		{
			genType = field.PURE_RANDOM_GENERATION
		}
	case "PERLIN_NOISE":
		{
			genType = field.PERLIN_NOISE_GENERATION
		}
	default:
		{
			fmt.Printf("Unknown generator: %v\n", *genTypeFlag)
			os.Exit(1)
		}
	}

	field := field.GetNewField(X_SIZE, Y_SIZE, genType, MIN_POCKET_SIZE, SPAWN_SIZE)
	renderer := renderer.GetRenderer(TILE_SIZE, field, SCREEN_WIDTH, SCREEN_HEIGHT)

	game := game.GetNewGame(field, renderer)
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
