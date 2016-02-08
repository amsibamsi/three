// Package matrix does simple math on matrices with homogeneous coordinates.
package matrix

import (
	"math"
)

// A matrix holds 16 components, the 4 first elements make up the first row from left to right, and so on.
type Matrix [16]float64
