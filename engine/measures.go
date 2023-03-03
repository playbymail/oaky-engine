// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

// define values for all in-game measurements.
const (

	// CABLE is 1/10th of a nautical mile or 100 fathoms.
	CABLE = 185.3184 // meters

	// FATHOM is 1/1000 of a nautical mile.
	FATHOM = 1.853184 // meters

	// KNOT is the unit of velocity.
	//   We define the knot as 1 UK Admiralty nautical mile per hour.
	//     1 knot is equal to 30.8864 meters per minute.
	//   "Knot" is abbreviated as "kn."
	KNOT = 30.8864 // meters per minute

	// LEAGUE is three nautical miles at sea.
	LEAGUE = 5_559.552 // meters

	// NAUTICAL_MILE is the unit of distance.
	//   We define a nautical mile to be 1,853.184 meters, which is
	//   the UK Admiralty nautical mile.
	//
	//   We use the UK Admiralty nautical mile because
	//     1. It was in use during the time period covered by the game,
	//     2. We wanted to use a single unit for all navies.
	NAUTICAL_MILE = 1_853.184 // meters

)
