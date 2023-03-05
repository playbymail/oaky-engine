// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo

import (
	"fmt"
	"math"
)

// Point is an immutable cartesian coordinate.
type Point struct {
	X, Y float64
}

// Add a vector
func (p Point) Add(vectors ...Vector) Point {
	x, y := p.X, p.Y
	for _, v := range vectors {
		x, y = x+v.X, y+v.Y
	}
	return Point{X: x, Y: y}
}

// Bearing from p1 to p2 is the inclination of a line drawn between
// the two points measured as the counter-clockwise angle from the
// positive X-axis.
func (p Point) Bearing(p2 Point) Angle {
	if approx(p.Y) == approx(p2.Y) {
		// horizontal line
		if approx(p2.X) < approx(p.X) {
			// p2 is "left" of p1, so treat as a straight angle
			return Angle(math.Pi)
		}
		// otherwise, treat as the zero angle
		return Angle(0)
	}
	if approx(p.X) == approx(p2.X) {
		// vertical line
		if approx(p2.Y) < approx(p.Y) {
			// p2 is "below" p1, so treat as a reflex angle
			return Angle(3 * math.Pi / 2)
		}
		// otherwise, treat as a right angle
		return Angle(math.Pi * 2)
	}
	// the inclination is the tangent of the slope, tan(θ) = m.
	θ := math.Atan((p2.Y - p.Y) / (p2.X - p.X))
	if θ < 0 {
		// angles > 90deg will generate negative values and must
		// be normalized
		θ += math.Pi
	}
	// but the range depends on if the line is sloping up or down (as seen from p1).
	if p.Y > p2.Y {
		// sloping down, so normalize by adding another pi to it
		θ += math.Pi
	}
	return Angle(θ)
}

func (p Point) Distance(p2 Point) float64 {
	dx, dy := (p2.X - p.X), (p2.Y - p.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// String implements the Stringer interface
func (p Point) String() string {
	return fmt.Sprintf("(%.6f %.6f)", p.X, p.Y)
}
