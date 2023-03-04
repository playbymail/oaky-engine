// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import "math"

// DegreesToRadians converts degrees to radians.
// 0 degrees is north.
// Degrees (and Radians) increase clockwise.
// Clamps the results.
func DegreesToRadians(d int) float64 {
	for d < 0 {
		d += 360
	}
	for d >= 360 {
		d -= 360
	}
	return float64(d) * math.Pi / 180.0
}

func RadiansToDegrees(r float64) int {
	const twoPI = math.Pi + math.Pi
	for r < 0 {
		r = r + twoPI
	}
	for r >= twoPI {
		r = r - twoPI
	}
	return int(r * 180.0 / math.Pi)
}
