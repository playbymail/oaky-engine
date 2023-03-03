// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import "math"

// DistanceBetween returns the distance between two points
func DistanceBetween(from, to Coordinates) float64 {
	dx, dy := to.X-from.X, to.Y-from.Y
	return math.Sqrt(dx*dx + dy*dy)
}
