package utils

import "math"

const (
	DIRECTION_TOP = iota
	DIRECTION_RIGHT
	DIRECTION_BOTTOM
	DIRECTION_LEFT
)

type Vector2D struct {
	X int
	Y int
}

func (v Vector2D) IsAnyNegative() bool {
	return v.X < 0 || v.Y < 0
}

func (v Vector2D) WillItFit(high Vector2D) bool {
	return v.X > -1 && v.X < high.X && v.Y > -1 && v.Y < high.Y
}

func (v Vector2D) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v Vector2D) Scale(scalar float32) Vector2D {
	return Vector2D{X: int(float32(v.X) * scalar), Y: int(float32(v.Y) * scalar)}
}

func (v Vector2D) Absolute() Vector2D {
	return Vector2D{
		X: int(math.Abs(float64(v.X))),
		Y: int(math.Abs(float64(v.X))),
	}
}

func (v Vector2D) GetSmallestCoordinate() int {
	if v.X > v.Y {
		return v.Y
	} else {
		return v.X
	}
}

func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 Vector2D) Divide(v2 Vector2D) Vector2D {
	return Vector2D{X: v1.X / v2.X, Y: v1.Y / v2.Y}
}

func (v1 Vector2D) Equals(v2 Vector2D) bool {
	return v1.X == v2.X && v1.Y == v2.Y
}

func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	dx := v2.X - v1.X
	dy := v2.Y - v1.Y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
func Contains(tiles *[]Vector2D, vec Vector2D) bool {
	for _, v := range *tiles {
		if v.Equals(vec) {
			return true
		}
	}
	return false
}

func AddUnique(tiles *[]Vector2D, vec Vector2D) {
	if !Contains(tiles, vec) {
		*tiles = append(*tiles, vec)
	}
}
