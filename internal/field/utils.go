package field

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/thneutral/go-trpg-game/internal/character"
	"github.com/thneutral/go-trpg-game/internal/utils"
)

const (
	PURE_RANDOM_GENERATION = iota
	PERLIN_NOISE_GENERATION
	OPEN_SIMPLEX_NOISE_GENERATION
)

var (
	BLACK_PURE = color.Black
	WHITE_PURE = color.White

	RED_PURE   = color.RGBA{R: 255}
	GREEN_PURE = color.RGBA{G: 255}
	BLUE_PURE  = color.RGBA{B: 255}

	RED_BASE   = color.RGBA{R: 0xad, G: 0x11, B: 0x11}
	GREEN_BASE = color.RGBA{R: 0x55, G: 0xc2, B: 0x72}

	THAT_BLUE = color.RGBA{R: 0x49, G: 0xE0, B: 0xE4}
)

func GetNewField(xsize, ysize int, generationType int, minPocketSize int, spawnSize int) *Field {
	// var seed int64 = 1730659138505970900
	seed := time.Now().UnixNano()

	field := &Field{
		XSize:         xsize,
		YSize:         ysize,
		MinPocketSize: minPocketSize,
		SpawnSize:     spawnSize,
		Seed:          seed,
	}

	switch generationType {
	case PURE_RANDOM_GENERATION:
		{
			field.GenerateMapPureRandom(seed, 75)
		}
	case PERLIN_NOISE_GENERATION:
		{
			field.GenerateMapPerlinNoise(seed, 1.5, 2, 3, 10)
		}
	case OPEN_SIMPLEX_NOISE_GENERATION:
		{
			field.GenerateMapOpenSimplexNoise(seed, 8)
		}
	default:
		{
			fmt.Printf("Incorrect generation type %v\n", generationType)
			os.Exit(1)
		}
	}

	field.ConnectClosedRegions()

	field.GenerateCharacters([]character.Character{{Mobility: 10, IsMine: true, Type: character.MELEE}})

	fmt.Printf("Generation took %v milliseconds", time.Now().UnixMilli()-(seed/(1000*1000)))

	return field
}

type Tile struct {
	Character *character.Character
	Color     color.Color
}

type Field struct {
	Seed          int64
	SpawnSize     int
	MinPocketSize int
	XSize         int
	YSize         int
	Tiles         []Tile
}

func (field *Field) GetColor(pos utils.Vector2D) color.Color {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	return field.Tiles[index].Color
}

func (field *Field) GetCharacter(pos utils.Vector2D) *character.Character {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	return field.Tiles[index].Character
}

func (field *Field) SetColor(pos utils.Vector2D, color color.Color) {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	field.Tiles[index].Color = color
}

func (field *Field) SetCharacter(pos utils.Vector2D, character *character.Character) {
	index := utils.TwoDimToOneDim(int(pos.X), int(pos.Y), field.XSize)
	field.Tiles[index].Character = character
}
