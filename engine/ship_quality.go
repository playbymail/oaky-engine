// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"bytes"
	"fmt"
)

type ShipQuality int

const (
	HOOKER  ShipQuality = -1
	AVERAGE ShipQuality = 0
	SMART   ShipQuality = 1
)

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
	}
	return fmt.Errorf("invalid code")
}
