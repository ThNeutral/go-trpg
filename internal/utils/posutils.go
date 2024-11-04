package utils

import (
	"fmt"
	"os"
)

func TwoDimToOneDim(x, y, xMax int) int {
	return y*xMax + x
}

func OneDimToTwoDim(index, xMax int) (int, int) {
	y := index / xMax
	x := index % xMax
	return x, y
}

func GetDirection(dir int) Vector2D {
	switch dir {
	case DIRECTION_TOP:
		{
			return Vector2D{X: 0, Y: -1}
		}
	case DIRECTION_RIGHT:
		{
			return Vector2D{X: 1, Y: 0}
		}
	case DIRECTION_BOTTOM:
		{
			return Vector2D{X: 0, Y: 1}
		}
	case DIRECTION_LEFT:
		{
			return Vector2D{X: -1, Y: 0}
		}
	default:
		{
			fmt.Printf("Incorrect direction %v\n", dir)
			os.Exit(1)
			return Vector2D{}
		}
	}
}

// Top, Right, Bottom, Left
func GetAllDirections() [4]Vector2D {
	return [4]Vector2D{GetDirection(DIRECTION_TOP), GetDirection(DIRECTION_RIGHT), GetDirection(DIRECTION_BOTTOM), GetDirection(DIRECTION_LEFT)}
}
