// Package vector does simple math on vectors with homogeneous coordinates.
package vector

import (
	"math"
)

// A vector holds 4 components: x, y, z and w in this order.
type Vector [4]float64
