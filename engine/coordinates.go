// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"bytes"
	"fmt"
	"math"
)

// NORTH and SOUTH are on the Y axis.
// NORTH is a positive Y value, SOUTH is negative.
//
// EAST and WEST are on the X axis.
// EAST is a positive X value, WEST is negative.
//
// Bearing is the angle between two points measured clockwise from NORTH.
//
// Relative bearing is measured clockwise from the ship's centerline.

// Coordinates are x and y (internally).
// At some point, we'll output them as longitude and latitude.
type Coordinates struct {
	X float64
	Y float64
}

// String implements the Stringer interface {
func (a Coordinates) String() string {
	return fmt.Sprintf("(%d %d)", int(math.Floor(a.X)), int(math.Floor(a.Y)))
}

// MarshalJSON implements the json.Marshaler interface.
func (a Coordinates) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("(%g %g)", a.X, a.Y)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Coordinates) UnmarshalJSON(b []byte) error {
	if n, err := fmt.Fscanf(bytes.NewReader(b), "(%g %g)", &a.X, &a.Y); n != 2 || err != nil {
		return fmt.Errorf("invalid coordinates")
	}
	return nil
}
