// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import "math"

// Bearing is the angle (measured clockwise from the heading) between two points.
//
// Absolute Bearing is the bearing when the heading is "true north."
// Relative Bearing is the bearing when the heading is the alignment of the ship.

// AbsoluteBearing returns the absolute bearing and from point A to point B.
func AbsoluteBearing(from, to Coordinates) float64 {
	bearing := math.Atan2((to.X - from.X), (to.Y - from.Y))
	for bearing < 0 {
		bearing += 2 * math.Pi
	}
	return bearing
}
