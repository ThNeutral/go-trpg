package field

import (
	"fmt"

	"math/rand"

	"github.com/aquilax/go-perlin"
	"github.com/ojrac/opensimplex-go"
)

func (field *Field) GenerateMapOpenSimplexNoise(seed int64, scalar float64) {
	fmt.Printf("Generating OpenSimplex noise map.\nCurrent seed: %v.\n", seed)
	generator := opensimplex.New(seed)
	arr := make([]Tile, 0)
	for y := range field.YSize {
		for x := range field.XSize {
			var t Tile
			if generator.Eval2(float64(x)/scalar, float64(y)/scalar) > 0.0 {
				t.Color = WHITE_PURE
			} else {
				t.Color = BLACK_PURE
			}
			arr = append(arr, t)
		}
	}
	field.Tiles = arr
}

func (field *Field) GenerateMapPerlinNoise(seed int64, alpha float64, beta float64, n int32, scalar float64) {
	fmt.Printf("Generating Perlin noise map.\nCurrent seed: %v.\n", seed)
	generator := perlin.NewPerlin(alpha, beta, n, seed)
	arr := make([]Tile, 0)
	for y := range field.YSize {
		for x := range field.XSize {
			var t Tile
			if generator.Noise2D(float64(x)/scalar, float64(y)/scalar) > 0.0 {
				t.Color = WHITE_PURE
			} else {
				t.Color = BLACK_PURE
			}
			arr = append(arr, t)
		}
	}
	field.Tiles = arr
}

func (field *Field) GenerateMapPureRandom(seed int64, chance int) {
	fmt.Printf("Generating pure random map.\nCurrent seed: %v.\n", seed)
	generator := rand.New(rand.NewSource(seed))
	arr := make([]Tile, 0)
	for range field.YSize {
		for range field.XSize {
			var t Tile
			if generator.Int31n(100) < int32(chance) {
				t.Color = WHITE_PURE
			} else {
				t.Color = BLACK_PURE
			}
			arr = append(arr, t)
		}
	}
	field.Tiles = arr
}
