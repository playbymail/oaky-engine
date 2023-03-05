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

// Coordinates are immutable Cartesian coordinates.
// At some point, we'll output them as longitude and latitude.
type Coordinates struct {
	X float64
	Y float64
}

func (c Coordinates) Add(vectors ...Vector) Coordinates {
	x, y := c.X, c.Y
	for _, v := range vectors {
		x, y = x+v.X, y+v.Y
	}
	return Coordinates{X: x, Y: y}
}

// String implements the Stringer interface {
func (c Coordinates) String() string {
	return fmt.Sprintf("(%d %d)", int(math.Floor(c.X)), int(math.Floor(c.Y)))
}

// MarshalJSON implements the json.Marshaler interface.
func (c Coordinates) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("(%g %g)", c.X, c.Y)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Coordinates) UnmarshalJSON(b []byte) error {
	if n, err := fmt.Fscanf(bytes.NewReader(b), "(%g %g)", &c.X, &c.Y); n != 2 || err != nil {
		return fmt.Errorf("invalid coordinates")
	}
	return nil
}
