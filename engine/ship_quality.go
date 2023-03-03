// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"bytes"
	"fmt"
)

const (
	// HOOKER represents below-average quality build and setup.
	HOOKER ShipQuality = -1

	// AVERAGE represents an average quality build and setup.
	AVERAGE ShipQuality = 0

	// SMART represents an above average quality build and setup.
	SMART ShipQuality = 1
)

// ShipQuality is a measure of the craftsmanship and materials used
// to build the ship, plus factors like the ballast, hull, rigging,
// and even the ship's personal quirks.
type ShipQuality int

// String implements the Stringer interface.
func (sq ShipQuality) String() string {
	switch sq {
	case HOOKER:
		return "hooker"
	case AVERAGE:
		return "average"
	case SMART:
		return "smart"
	}
	panic(fmt.Sprintf("assert(shipQuality != %d)", sq))
}

// MarshalJSON implements the json.Marshaler interface.
func (sq ShipQuality) MarshalJSON() ([]byte, error) {
	switch sq {
	case HOOKER:
		return []byte("HOOKER"), nil
	case AVERAGE:
		return []byte("AVERAGE"), nil
	case SMART:
		return []byte("SMART"), nil
	}
	return nil, fmt.Errorf("invalid code")
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (sq *ShipQuality) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`"HOOKER"`)) {
		*sq = HOOKER
	} else if bytes.Equal(b, []byte(`"AVERAGE"`)) {
		*sq = AVERAGE
	} else if bytes.Equal(b, []byte(`"SMART"`)) {
		*sq = SMART
	} else {
		return fmt.Errorf("invalid ShipQuality")
	}
	return nil
}
