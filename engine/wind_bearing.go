// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import "fmt"

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
)

const (
	WB_ASTERN        WindBearing = 0
	WB_QUARTER_REACH             = 1
	WB_BROAD_REACH               = 2
	WB_BEATING                   = 3

	// sizeofWindBearing is used for sizing arrays.
	// It must be one more than the largest enum value.
	sizeofWindBearing = WB_BEATING + 1
)

// WindAngle is number of compass points the stern is from the wind.
type WindAngle int

type WindBearing int

func (wb WindBearing) String() string {
	switch wb {
	case WB_ASTERN:
		return "ASTERN"
	case WB_QUARTER_REACH:
		return "QUARTER_REACH"
	case WB_BROAD_REACH:
		return "BROAD_REACH"
	case WB_BEATING:
		return "BEATING"
	}
	panic(fmt.Sprintf("assert(WindBearing != %d)", wb))
}
