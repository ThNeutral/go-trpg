package field

import (
	"math/rand"

	"github.com/thneutral/go-trpg-game/internal/character"
	"github.com/thneutral/go-trpg-game/internal/utils"
)

func (field *Field) GetSpawnTiles(isAlly bool) []utils.Vector2D {
	generator := rand.New(rand.NewSource(field.Seed))
	var tiles []utils.Vector2D
	var baseCell utils.Vector2D
	color := GREEN_BASE
	if !isAlly {
		color = RED_BASE
		baseCell = utils.Vector2D{X: field.XSize - 1, Y: field.YSize - 1}
	}
	for field.GetColor(baseCell) == BLACK_PURE {
		newCell := baseCell.Add(utils.GetDirection(generator.Intn(4)))
		if newCell.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
			baseCell = newCell
		}
	}
	field.SetColor(baseCell, color)
	for range field.SpawnSize {
	}
	return tiles
}

func (field *Field) GenerateCharacters(characters []character.Character) {
	field.GetSpawnTiles(true)
	field.GetSpawnTiles(false)
}
