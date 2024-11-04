package field

import (
	"container/list"
	"image/color"

	"github.com/thneutral/go-trpg-game/internal/utils"
)

func (field *Field) CreateSinglePath(origin utils.Vector2D, end utils.Vector2D, tilesToUse []utils.Vector2D) []utils.Vector2D {
	tileSet := make(map[utils.Vector2D]bool)
	visited := make(map[utils.Vector2D]bool)
	for _, tile := range tilesToUse {
		tileSet[tile] = true
		visited[tile] = false
	}
	if !tileSet[origin] || !tileSet[end] {
		return []utils.Vector2D{}
	}

	queue := list.New()
	queue.PushBack([]utils.Vector2D{origin})
	path := make([]utils.Vector2D, 0)
	for queue.Len() > 0 {
		currentPath := queue.Remove(queue.Front()).([]utils.Vector2D)
		currentPosition := currentPath[len(currentPath)-1]

		if currentPosition.Equals(end) {
			path = currentPath
			break
		}

		for _, dir := range utils.GetAllDirections() {
			next := currentPosition.Add(dir)

			if !next.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
				continue
			}
			if visited[next] {
				continue
			}
			if !tileSet[next] {
				continue
			}
			nextPath := append([]utils.Vector2D{}, currentPath...)
			nextPath = append(nextPath, next)
			queue.PushBack(nextPath)
		}
	}
	if len(path) == 0 {
		return []utils.Vector2D{}
	}
	for i, pos := range path {
		if i == 0 || i == len(path)-1 {
			field.SetColor(pos, RED_PURE)
			continue
		}
		field.SetColor(pos, THAT_BLUE)
	}
	return path
}

func (field *Field) CreateAllPaths(origin utils.Vector2D, len int) []utils.Vector2D {
	tiles := &[]utils.Vector2D{}
	field.computePathsInternal(origin, len, 0, THAT_BLUE, tiles)
	return *tiles
}

func (field *Field) DeleteAllPaths(origin utils.Vector2D, len int) {
	field.computePathsInternal(origin, len, 0, WHITE_PURE, nil)
}

func (field *Field) computePathsInternal(origin utils.Vector2D, len int, depth int, color color.Color, tiles *[]utils.Vector2D) {
	if depth > len {
		return
	}
	field.SetColor(origin, color)
	for _, dirs := range utils.GetAllDirections() {
		new := origin.Add(dirs)
		if !new.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
			continue
		}
		if field.GetColor(new) == BLACK_PURE {
			continue
		}
		if tiles != nil {
			utils.AddUnique(tiles, new)
		}
		field.computePathsInternal(new, len, depth+1, color, tiles)
	}
}
