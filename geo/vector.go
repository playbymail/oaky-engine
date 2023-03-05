// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo

import (
	"fmt"
	"math"
)

// Vector is and length (magnitude) and direction.
type Vector struct {
	// represent as a slope
	X, Y float64
	// represent as a length and direction
	L float64
	A Angle
}

func VectorFromPoint(x, y float64) Vector {
	return VectorFromPoints(Point{0, 0}, Point{x, y})
}

func VectorFromPoints(p1, p2 Point) Vector {
	// sin(A) = y/z <=> y = z*sin(A)
	// cos(A) = x/z <=> x = z*cos(A)
	a, z := p1.Bearing(p2), p1.Distance(p2)
	return Vector{
		X: p2.X - p1.X, //z * math.Cos(float64(a)),
		Y: p2.Y - p1.Y, //z * math.Sin(float64(a)),
		L: z,
		A: a,
	}
}

// Add vectors
func (v Vector) Add(vectors ...Vector) Vector {
	x, y := v.X, v.Y
	for _, vec := range vectors {
		x, y = x+vec.X, y+vec.Y
	}
	return VectorFromPoint(x, y)
}

// Angle returns the angle of the vector.
func (v Vector) Angle() Angle {
	return v.A
}

// Length and magnitude are the same.
func (v Vector) Length() float64 {
	return v.L
}

// RelativeAngle translates an absolute angle to an angle relative to the vector.
func (v Vector) RelativeAngle(a Angle) Angle {
	return Angle(2*math.Pi - (v.A - a)).Normalize()
}

// String is a stringer utility function for vectors
func (v Vector) String() string {
	return fmt.Sprintf("[%.6f,%.6f]", v.X, v.Y)
}
