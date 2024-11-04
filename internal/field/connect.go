package field

import (
	"math"
	"slices"

	"github.com/thneutral/go-trpg-game/internal/utils"
)

func (field *Field) ConnectClosedRegions() {
	clusters := [][]utils.Vector2D{}
	for len(clusters) != 1 {
		clusters = field.countClusters()
		field.connectClusters(clusters)
	}
}

func (field *Field) countClusters() [][]utils.Vector2D {
	clusters := [][]utils.Vector2D{}
	visited := make(map[utils.Vector2D]bool)
	for y := range field.YSize {
		for x := range field.XSize {
			cell := utils.Vector2D{X: x, Y: y}
			if field.GetColor(cell) == WHITE_PURE && !visited[cell] {
				clusters = append(clusters, field.floodFill(cell, visited))
			}
		}
	}
	return clusters
}

func (field *Field) floodFill(start utils.Vector2D, visited map[utils.Vector2D]bool) []utils.Vector2D {
	var stack []utils.Vector2D
	pocket := []utils.Vector2D{}
	stack = append(stack, start)

	for len(stack) > 0 {
		cell := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cell.X < 0 || cell.Y < 0 || cell.X >= field.XSize || cell.Y >= field.YSize || visited[cell] || field.GetColor(cell) == BLACK_PURE {
			continue
		}

		visited[cell] = true
		pocket = append(pocket, cell)

		for _, dir := range utils.GetAllDirections() {
			stack = append(stack, cell.Add(dir))
		}
	}

	return pocket
}

func (field *Field) connectClusters(clusters [][]utils.Vector2D) {
	type Edge struct {
		From     utils.Vector2D
		To       utils.Vector2D
		Distance float32
	}
	edges := []Edge{}
	for _, clusterA := range clusters {
		var edge Edge
		edge.Distance = float32(math.MaxFloat32)
		for _, clusterB := range clusters {
			if slices.Equal(clusterA, clusterB) {
				continue
			}
			if newMin, from, to := field.minimalLength(clusterA, clusterB); newMin < edge.Distance {
				edge.To = to
				edge.From = from
				edge.Distance = newMin
			}
		}
		edges = append(edges, edge)
	}
	for _, edge := range edges {
		field.createPath(edge.From, edge.To)
	}
}

func (field *Field) createPath(cellA, cellB utils.Vector2D) {
	dx := cellB.X - cellA.X
	dy := cellB.Y - cellA.Y

	absDx := int(math.Abs(float64(dx)))
	absDy := int(math.Abs(float64(dy)))

	stepX := 1
	if dx < 0 {
		stepX = -1
	}
	stepY := 1
	if dy < 0 {
		stepY = -1
	}

	if absDx > absDy {
		D := 2*absDy - absDx
		y := cellA.Y

		for x := cellA.X; x != cellB.X; x += stepX {
			cell := utils.Vector2D{X: x, Y: y}
			if !cell.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
				continue
			}
			field.SetColor(cell, WHITE_PURE)

			for _, dir := range utils.GetAllDirections() {
				newCell := cell.Add(dir)
				if !newCell.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
					continue
				}
				field.SetColor(newCell, WHITE_PURE)
			}

			if D > 0 {
				y += stepY
				D -= 2 * absDx
			}
			D += 2 * absDy
		}
	} else {
		D := 2*absDx - absDy
		x := cellA.X

		for y := cellA.Y; y != cellB.Y; y += stepY {
			cell := utils.Vector2D{X: x, Y: y}
			if !cell.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
				continue
			}
			field.SetColor(cell, WHITE_PURE)

			for _, dir := range utils.GetAllDirections() {
				newCell := cell.Add(dir)
				if !newCell.WillItFit(utils.Vector2D{X: field.XSize, Y: field.YSize}) {
					continue
				}
				field.SetColor(newCell, WHITE_PURE)
			}

			if D > 0 {
				x += stepX
				D -= 2 * absDy
			}
			D += 2 * absDx
		}
	}
}

func (field *Field) minimalLength(from []utils.Vector2D, to []utils.Vector2D) (float32, utils.Vector2D, utils.Vector2D) {
	min := float32(math.MaxFloat32)
	var fromCellReturn utils.Vector2D
	var toCellReturn utils.Vector2D
	for _, cellFrom := range from {
		for _, cellTo := range to {
			if dist := float32(cellFrom.Distance(cellTo)); dist < min {
				fromCellReturn = cellFrom
				toCellReturn = cellTo
				min = dist
			}
		}
	}
	return min, fromCellReturn, toCellReturn
}
