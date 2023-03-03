// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"fmt"
	"math"
)

// Vector is a two-dimensional vector
type Vector struct {
	X, Y float64
}

// Add returns the sum of two vectors
func (v Vector) Add(w Vector) Vector {
	return Vector{X: v.X + w.X, Y: v.Y + w.Y}
}

// Equals returns true only if the two vectors have the same representation.
// (In other words, if they look the same when we print them out.)
func (v Vector) Equals(w Vector) bool {
	return v.String() == w.String()
}

// Length returns the magnitude of a vector
func (v Vector) Length() float64 {
	return v.Magnitude()
}

// Magnitude returns the magnitude of a vector.
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalize returns a normalized vector.
func (v Vector) Normalize() Vector {
	magnitude := v.Magnitude()
	if magnitude == 0.0 {
		return Vector{}
	}
	return Vector{X: v.X / magnitude, Y: v.Y / magnitude}
}

// RotateDegrees returns a vector rotated anti-clockwise by β degrees
func (v Vector) RotateDegrees(β float64) Vector {
	return v.RotateRadians(β * math.Pi / 180.0)
}

// RotateRadians returns a vector rotated anti-clockwise by β radians
func (v Vector) RotateRadians(β float64) Vector {
	sinβ, cosβ := math.Sin(β), math.Cos(β)
	return Vector{X: cosβ*v.X - sinβ*v.Y, Y: sinβ*v.X + cosβ*v.Y}
}

// Scale returns a vector multiplied by a scalar.
func (v Vector) Scale(s float64) Vector {
	return Vector{X: v.X * s, Y: v.Y * s}
}

// String implements the Stringer interface.
func (v Vector) String() string {
	return fmt.Sprintf("<%.4f %.4f>", v.X, v.Y)
}
