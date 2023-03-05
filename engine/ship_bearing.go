// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

// really bearing from any vessel that has a heading

// BearingTo returns the relative bearing from the ship to
// another object based on the ship's current heading.
// Bearing is measured in radians and clockwise from the heading.
func (s Ship) BearingTo(object Coordinates) float64 {
	// move both ship and object to origin
	moveVector := Vector{X: -s.Coordinates.X, Y: -s.Coordinates.Y}
	//fmt.Printf("ship %-10s: object %-10s: move %-10s\n", s.Coordinates, object, moveVector)
	sa, so := s.Coordinates.Add(moveVector), object.Add(moveVector)
	//fmt.Printf("ship %-10s: object %-10s: move %-10s\n", sa, so, moveVector)
	// todo: rotate ship to heading of true north
	rotateVector := Vector{so.X, so.Y}.RotateRadians(s.Heading)
	//fmt.Printf("ship %-10s: object %-10s: rotate %-10s\n", sa, so, rotateVector)
	// return bearing
	bearing := AbsoluteBearing(sa, Coordinates{rotateVector.X, rotateVector.Y})
	//fmt.Printf("ship %-10s: object %-10s: bear %8.04f degrees %8d\n", sa, so, bearing, RadiansToDegrees(bearing))
	return bearing
}
