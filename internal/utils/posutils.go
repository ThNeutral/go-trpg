package utils

func TwoDimToOneDim(x, y, xMax int) int {
	return y*xMax + x
}

func OneDimToTwoDim(index, xMax int) (int, int) {
	y := index / xMax
	x := index % xMax
	return x, y
}

type Vector2D struct {
	X float32
	Y float32
}
