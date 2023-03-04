// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

// See https://en.wikipedia.org/wiki/Rating_system_of_the_Royal_Navy
type ShipRate int

const (
	NOT_RATED ShipRate = iota
	SHIPS_BOAT

	// CUTTER is single masted vessel.
	// See https://en.wikipedia.org/wiki/Cutter_(boat)
	CUTTER

	// LUGGER is 2 or 3 masted vessel with lug sails.
	// Used for smuggling and piracy.
	// See https://en.wikipedia.org/wiki/Lugger
	LUGGER

	// See https://en.wikipedia.org/wiki/Xebec
	XEBEC

	// BOMB is a ketch with two masts.
	// See https://en.wikipedia.org/wiki/Bomb_vessel
	BOMB_VESSEL

	// BRIG is 2 masted vessel.
	// Seehttps://en.wikipedia.org/wiki/Brig
	BRIG

	// See https://en.wikipedia.org/wiki/Sloop-of-war
	SLOOP

	// CORVETTE is larger than a sloop.
	// See https://en.wikipedia.org/wiki/Corvette
	CORVETTE

	FRIGATE_6TH_RATE
	FRIGATE_5TH_RATE
	FRIGATE_4TH_RATE
	SOL_4TH_RATE
	SOL_3RD_RATE
	SOL_2ND_RATE
	SOL_1ST_RATE

	// sizeofShipRate is used for sizing arrays.
	// It must be the last value defined for the enums.
	SizeofShipRate
)
