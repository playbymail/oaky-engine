// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package compass_rose

import "fmt"

// Bearing does not belong here
func (p Point) Bearing() string {
	switch p {
	case N:
		return "Dead Ahead"
	case NbE:
		return "1 point off the Starboard Bow"
	case NNE:
		return "2 points off the Starboard Bow"
	case NEbN:
		return "3 points off the Starboard Bow"
	case NE:
		return "4 points off the Starboard Bow"
	case NEbE:
		return "5 points off the Starboard Bow"
	case ENE:
		return "6 points off the Starboard Bow"
	case EbN:
		return "7 points off the Starboard Bow"
	case E:
		return "Off the Starboard Side"
	case EbS:
		return "7 points off the Starboard Stern"
	case ESE:
		return "6 points off the Starboard Stern"
	case SEbE:
		return "5 points off the Starboard Stern"
	case SE:
		return "4 points off the Starboard Stern"
	case SEbS:
		return "3 points off the Starboard Stern"
	case SSE:
		return "2 points off the Starboard Stern"
	case SbE:
		return "1 point off the Starboard Stern"
	case S:
		return "Astern"
	case SbW:
		return "1 point off the Port Stern"
	case SSW:
		return "2 points off the Port Stern"
	case SWbS:
		return "3 points off the Port Stern"
	case SW:
		return "4 points off the Port Stern"
	case SWbW:
		return "5 points off the Port Stern"
	case WSW:
		return "6 points off the Port Stern"
	case WbS:
		return "7 points off the Port Stern"
	case W:
		return "Off the Port Side"
	case WbN:
		return "7 points off the Port Bow"
	case WNW:
		return "6 points off the Port Bow"
	case NWbW:
		return "5 points off the Port Bow"
	case NW:
		return "4 points off the Port Bow"
	case NWbN:
		return "3 points off the Port Bow"
	case NNW:
		return "2 points off the Port Bow"
	case NbW:
		return "1 point off the Port Bow"
	}
	panic(fmt.Sprintf("assert(point != %d)", p))
}
