// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package compass_rose

import "fmt"

// Point is a compass rose point.
// There are 32 points on the compass and each corresponds to an absolute bearing.
type Point int

// enums for Point
const (
	N Point = iota
	NbE
	NNE
	NEbN
	NE
	NEbE
	ENE
	EbN
	E
	EbS
	ESE
	SEbE
	SE
	SEbS
	SSE
	SbE
	S
	SbW
	SSW
	SWbS
	SW
	SWbW
	WSW
	WbS
	W
	WbN
	WNW
	NWbW
	NW
	NWbN
	NNW
	NbW
)

// String implements the Stringer interface
func (p Point) String() string {
	switch p {
	case N:
		return "N"
	case NbW:
		return "NbW"
	case NNW:
		return "NNW"
	case NWbN:
		return "NWbN"
	case NW:
		return "NW"
	case NWbW:
		return "NWbW"
	case WNW:
		return "WNW"
	case WbN:
		return "WbN"
	case W:
		return "W"
	case WbS:
		return "WbS"
	case WSW:
		return "WSW"
	case SWbW:
		return "SWbW"
	case SW:
		return "SW"
	case SWbS:
		return "SWbS"
	case SSW:
		return "SSW"
	case SbW:
		return "SbW"
	case S:
		return "S"
	case SbE:
		return "SbE"
	case SSE:
		return "SSE"
	case SEbS:
		return "SEbS"
	case SE:
		return "SE"
	case SEbE:
		return "SEbE"
	case ESE:
		return "ESE"
	case EbS:
		return "EbS"
	case E:
		return "E"
	case EbN:
		return "EbN"
	case ENE:
		return "ENE"
	case NEbE:
		return "NEbE"
	case NE:
		return "NE"
	case NEbN:
		return "NEbN"
	case NNE:
		return "NNE"
	case NbE:
		return "NbE"
	}
	panic(fmt.Sprintf("assert(point != %d)", p))
}

// Long description of the points.
func (p Point) Long() string {
	switch p {
	case N:
		return "North"
	case NbE:
		return "North by East"
	case NNE:
		return "North NorthEast"
	case NEbN:
		return "NorthEast by North"
	case NE:
		return "NorthEast"
	case NEbE:
		return "NorthEast by East"
	case ENE:
		return "East NorthEast"
	case EbN:
		return "East by North"
	case E:
		return "East"
	case EbS:
		return "East by South"
	case ESE:
		return "East SouthEast"
	case SEbE:
		return "SouthEast by East"
	case SE:
		return "SouthEast"
	case SEbS:
		return "SouthEast by South"
	case SSE:
		return "South SouthEast"
	case SbE:
		return "South by East"
	case S:
		return "South"
	case SbW:
		return "South by West"
	case SSW:
		return "South SouthWest"
	case SWbS:
		return "SouthWest by South"
	case SW:
		return "SouthWest"
	case SWbW:
		return "SouthWest by West"
	case WSW:
		return "West SouthWest"
	case WbS:
		return "West by South"
	case W:
		return "West"
	case WbN:
		return "West by North"
	case WNW:
		return "West NorthWest"
	case NWbW:
		return "NorthWest by West"
	case NW:
		return "NorthWest"
	case NWbN:
		return "NorthWest by North"
	case NNW:
		return "North NorthWest"
	case NbW:
		return "North by West"
	}
	panic(fmt.Sprintf("assert(point != %d)", p))
}
