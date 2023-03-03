// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import "fmt"

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

// String implements the Stringer interface.
func (v Vector) String() string {
	return fmt.Sprintf("<%.4f %.4f>", v.X, v.Y)
}
