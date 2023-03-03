// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

/*
 */

type SailSetting int

const (
	// Ship will drift unless anchored.
	NO_SAIL SailSetting = iota

	// Generally headsails and driver.
	// Barely enough sail to give the ship headway,
	// allow the rudder to bite, and allow for fine maneuvering.
	MINIMUM_SAIL

	// Generally reefed topsails.
	// A good setting for battle, as it provides a stable
	// gunnery platform and allows for fine maneuvering.
	FIGHTING_SAIL

	// Generally mainsails and topsails. Easy to manage.
	ALL_PLAIN_SAIL

	// Topgallants and royals.
	// Fast, difficult to manage, lacks easy maneuverability.
	FULL_SAIL

	// Courses, studding sails, gaffsails, staysails, and
	// other weird creations of the sailmaker’s art.
	// Hard to handle, very fast, with very little maneuverability.
	EXTRA_SAIL
)
