package utils

import (
	"slices"
)

func AssertIntOneOf(expected []int, actual int, message string) {
	if !slices.Contains(expected, actual) {
		panic(message)
	}
}
