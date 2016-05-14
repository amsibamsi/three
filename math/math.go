package math

import (
	"math"
)

// Absi returns the absolute value of an integer.
func Absi(i int) int {
	m := i >> 31
	return (m ^ i) - m
}

// Round returns the nearest integer for a float.
func Round(f float64) int {
	return int(math.Floor(f + 0.5))
}

// Maxi returns the maximum of two integers.
func Maxi(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// Mini returns the minimum of two integers.
func Mini(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
