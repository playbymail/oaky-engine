// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package ship

import "github.com/playbymail/oaky-engine/engine"

// really bearing from any vessel that has a heading

// BearingTo returns the relative bearing from the ship to
// another object based on the ship's current heading.
// Bearing is measured in radians and clockwise from the heading.
func (s Ship) BearingTo(object engine.Coordinates) float64 {
	// move both ship and object to origin
	moveVector := engine.Vector{X: -s.coordinates.X, Y: -s.coordinates.Y}
	//fmt.Printf("ship %-10s: object %-10s: move %-10s\n", s.Coordinates, object, moveVector)
	sa, so := s.coordinates.Add(moveVector), object.Add(moveVector)
	//fmt.Printf("ship %-10s: object %-10s: move %-10s\n", sa, so, moveVector)
	// todo: rotate ship to heading of true north
	rotateVector := engine.Vector{so.X, so.Y}.RotateRadians(s.Heading())
	//fmt.Printf("ship %-10s: object %-10s: rotate %-10s\n", sa, so, rotateVector)
	// return bearing
	bearing := engine.AbsoluteBearing(sa, engine.Coordinates{rotateVector.X, rotateVector.Y})
	//fmt.Printf("ship %-10s: object %-10s: bear %8.04f degrees %8d\n", sa, so, bearing, RadiansToDegrees(bearing))
	return bearing
}
