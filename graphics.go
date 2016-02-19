// Render 3D graphics
package graphics

import ()

// Vector with homogeneous coordinates.
// Holds 4 components: x, y, z and w in this order.
type Vec4 [4]float64

// Matrix with homogeneous coordinates.
// Holds 16 components, the 4 first elements make up the first row from left to right, and so on.
type Mat4 [16]float64
