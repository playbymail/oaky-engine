// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"bytes"
	"fmt"
)

const (
	// Angle to the wind affects both speed and handling. The angle is measured
	// as the number of compass points the stern is from the wind.
	//
	// Ahead           the wind is DIRECTLY ahead.
	// Astern          the wind is DIRECTLY astern.
	// Quarter Reach   the wind is from 1 to 3 points from the stern.
	//                 This is the fastest possible position, since all
	//                 sails are in a position to catch the wind.
	// Broad Reach     the wind is 4 or 5 points from the stern.
	//                 Fairly slow, since not all the sails will draw,
	//                 and the wind is working directly against the
	//                 action of the keel against the water.
	// Beating         the stern is 6 or 7 points from the wind.
	//                 (Note that most ships can not sail with their sterns 7 points from the wind).
	//                 The slowest possible position. At this position, the
	//                 sails cease to act as a kite and begin to act as an airfoil.
	//                 The passage of the wind over the sails creates a vacuum
	//                 (actually a low pressure area) in front of the ship that
	//                 "sucks" the vessel forward. At this angle to the wind, a
	//                 ship may tack (pass its bow across the wind in order to
	//                 change direction), wear (change direction by passing its
	//                 stern across the wind), heave-to (stop and drift without
	//                 taking off all sails), and back sails (spill wind from the
	//                 sails and slow the ship, giving very exact control of movement).

	WIND_ASTERN      WindAngle = 0  // The wind is DIRECTLY astern.
	QUARTER_REACH_P1           = 1  // Quarter Reach, 1 point to port.
	QUARTER_REACH_S1           = -1 // Quarter Reach, 1 point to starboard.
	QUARTER_REACH_P2           = 2  // Quarter Reach, 2 points to port.
	QUARTER_REACH_S2           = -2 // Quarter Reach, 2 points to starboard.
	QUARTER_REACH_P3           = 3  // Quarter Reach, 3 points to port.
	QUARTER_REACH_S3           = -3 // Quarter Reach, 3 points to starboard.
	BROAD_REACH_P4             = 4  // Broad Reach, 4 points to port.
	BROAD_REACH_S4             = -4 // Broad Reach, 4 points to starboard.
	BROAD_REACH_P5             = 5  // Broad Reach, 5 points to port.
	BROAD_REACH_S5             = -5 // Broad Reach, 5 points to starboard.
	BEATING_P6                 = 6  // Beating, 6 points to port.
	BEATING_S6                 = -6 // Beating, 6 points to starboard.
	BEATING_P7                 = 7  // Beating, 7 points to port.
	BEATING_S7                 = -7 // Beating, 7 points to starboard.
	WIND_AHEAD                 = 8  // The wind is DIRECTLY ahead.

	// Strength of the wind affects sailing ability.
	// The stronger the wind, the faster the vessel will move; but a
	// very strong wind can destroy a ship if it sets too much sail.
	//
	// These constants define the possible Wind Strengths.

	// Ship will not move unless it is moved under oars, or are being towed by boats.
	CALM WindSpeed = iota

	// Ship may "catch a wind" and move, while others will have to row, or be towed by boats.
	LIGHT_GUSTS

	// Light wind.
	LIGHT_WIND

	// Normal breeze.
	NORMAL_BREEZE

	// Ship may not use Sail Setting EXTRA_SAIL without losing masts.
	HEAVY_BREEZE

	// Ship may not use Sail Setting FULL_SAIL or higher without losing masts.
	// Lower gunports may be awash.
	GALE

	// Ship may not use Sail Setting ALL_PLAIN_SAIL or higher without losing masts.
	// Lower gunports may be awash.
	STORM

	// Ship may not use Sail Setting FIGHTING_SAIL or higher without losing masts.
	// No combat is allowed.
	HURRICANE
)

// WindAngle is number of compass points the stern is from the wind.
type WindAngle int

// WindSpeed is a measure of the current speed of the wind.
type WindSpeed int

// String implements the Stringer interface.
func (w WindSpeed) String() string {
	switch w {
	case CALM:
		return "CALM"
	case LIGHT_GUSTS:
		return "LIGHT_GUSTS"
	case LIGHT_WIND:
		return "LIGHT_WIND"
	case NORMAL_BREEZE:
		return "NORMAL_BREEZE"
	case HEAVY_BREEZE:
		return "HEAVY_BREEZE"
	case GALE:
		return "GALE"
	case STORM:
		return "STORM"
	case HURRICANE:
		return "HURRICANE"
	}
	panic(fmt.Sprintf("assert(WindSpeed != %d)", w))
}

// MarshalJSON implements the json.Marshaler interface.
func (w WindSpeed) MarshalJSON() ([]byte, error) {
	switch w {
	case CALM:
		return []byte("CALM"), nil
	case LIGHT_GUSTS:
		return []byte("LIGHT_GUSTS"), nil
	case LIGHT_WIND:
		return []byte("LIGHT_WIND"), nil
	case NORMAL_BREEZE:
		return []byte("NORMAL_BREEZE"), nil
	case HEAVY_BREEZE:
		return []byte("HEAVY_BREEZE"), nil
	case GALE:
		return []byte("GALE"), nil
	case STORM:
		return []byte("STORM"), nil
	case HURRICANE:
		return []byte("HURRICANE"), nil
	}
	return nil, fmt.Errorf("invalid WindSpeed")
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (w *WindSpeed) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`"CALM"`)) {
		*w = CALM
	} else if bytes.Equal(b, []byte(`"LIGHT_GUSTS"`)) {
		*w = LIGHT_GUSTS
	} else if bytes.Equal(b, []byte(`"LIGHT_WIND"`)) {
		*w = LIGHT_WIND
	} else if bytes.Equal(b, []byte(`"NORMAL_BREEZE"`)) {
		*w = NORMAL_BREEZE
	} else if bytes.Equal(b, []byte(`"HEAVY_BREEZE"`)) {
		*w = HEAVY_BREEZE
	} else if bytes.Equal(b, []byte(`"GALE"`)) {
		*w = GALE
	} else if bytes.Equal(b, []byte(`"STORM"`)) {
		*w = STORM
	} else if bytes.Equal(b, []byte(`"HURRICANE"`)) {
		*w = HURRICANE
	} else {
		return fmt.Errorf("invalid WindSpeed")
	}
	return nil
}
