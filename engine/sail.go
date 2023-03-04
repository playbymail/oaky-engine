// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

import (
	"bytes"
	"fmt"
)

// The quality and quantity of sail that a ship sets affects
// sailing ability. The more sail, the more speed; but the more
// speed, the less maneuverable a ship becomes. At the higher
// sail settings, a ship may zip past its Opponents and have to
// spend many turns trying to work its way back into position.
//
// These constants define the possible Sail Settings.
const (
	// Ship will drift unless anchored.
	NO_SAIL SailSetting = 0

	// Generally headsails and driver.
	// Barely enough sail to give the ship headway,
	// allow the rudder to bite, and allow for fine maneuvering.
	MINIMUM_SAIL = 1

	// Generally reefed topsails.
	// A good setting for battle, as it provides a stable
	// gunnery platform and allows for fine maneuvering.
	FIGHTING_SAIL = 2

	// Generally mainsails and topsails. Easy to manage.
	ALL_PLAIN_SAIL = 3

	// Topgallants and royals.
	// Fast, difficult to manage, lacks easy maneuverability.
	FULL_SAIL = 4

	// Courses, studding sails, gaffsails, staysails, and
	// other weird creations of the sailmaker's art.
	// Hard to handle, very fast, with very little maneuverability.
	EXTRA_SAIL = 5

	// sizeofSailSetting is used for sizing arrays.
	// It must be the last value defined for the enums.
	sizeofSailSetting = EXTRA_SAIL + 1
)

// SailSetting represents the quantity and quality of sails set.
type SailSetting int

// String implements the Stringer interface.
func (s SailSetting) String() string {
	switch s {
	case NO_SAIL:
		return "NO_SAIL"
	case MINIMUM_SAIL:
		return "MINIMUM_SAIL"
	case FIGHTING_SAIL:
		return "FIGHTING_SAIL"
	case ALL_PLAIN_SAIL:
		return "ALL_PLAIN_SAIL"
	case FULL_SAIL:
		return "FULL_SAIL"
	case EXTRA_SAIL:
		return "EXTRA_SAIL"
	}
	panic(fmt.Sprintf("assert(SailSetting != %d)", s))
}

// MarshalJSON implements the json.Marshaler interface.
func (s SailSetting) MarshalJSON() ([]byte, error) {
	switch s {
	case NO_SAIL:
		return []byte("NO_SAIL"), nil
	case MINIMUM_SAIL:
		return []byte("MINIMUM_SAIL"), nil
	case FIGHTING_SAIL:
		return []byte("FIGHTING_SAIL"), nil
	case ALL_PLAIN_SAIL:
		return []byte("ALL_PLAIN_SAIL"), nil
	case FULL_SAIL:
		return []byte("FULL_SAIL"), nil
	case EXTRA_SAIL:
		return []byte("EXTRA_SAIL"), nil
	}
	return nil, fmt.Errorf("invalid SailSetting")
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *SailSetting) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte(`"NO_SAIL"`)) {
		*s = NO_SAIL
	} else if bytes.Equal(b, []byte(`"MINIMUM_SAIL"`)) {
		*s = MINIMUM_SAIL
	} else if bytes.Equal(b, []byte(`"FIGHTING_SAIL"`)) {
		*s = FIGHTING_SAIL
	} else if bytes.Equal(b, []byte(`"ALL_PLAIN_SAIL"`)) {
		*s = ALL_PLAIN_SAIL
	} else if bytes.Equal(b, []byte(`"FULL_SAIL"`)) {
		*s = FULL_SAIL
	} else if bytes.Equal(b, []byte(`"EXTRA_SAIL"`)) {
		*s = EXTRA_SAIL
	} else {
		return fmt.Errorf("invalid SailSetting")
	}
	return nil
}

// AddSail increases the amount of sail, if possible.
// Ships may increase the amount of sail one level per turn.
//
// The updated value is returned.
func (s SailSetting) AddSail() SailSetting {
	switch s {
	case NO_SAIL:
		return MINIMUM_SAIL
	case MINIMUM_SAIL:
		return FIGHTING_SAIL
	case FIGHTING_SAIL:
		return ALL_PLAIN_SAIL
	case ALL_PLAIN_SAIL:
		return FULL_SAIL
	case FULL_SAIL:
		return EXTRA_SAIL
	case EXTRA_SAIL:
		return EXTRA_SAIL
	}
	panic(fmt.Sprintf("assert(SailSetting != %d)", s))
}

// ReduceSail decreases the amount of sail, if possible.
// Ships may decrease the amount of sail one level per turn.
//
// The updated value is returned.
func (s SailSetting) ReduceSail() SailSetting {
	switch s {
	case NO_SAIL:
		return NO_SAIL
	case MINIMUM_SAIL:
		return NO_SAIL
	case FIGHTING_SAIL:
		return MINIMUM_SAIL
	case ALL_PLAIN_SAIL:
		return FIGHTING_SAIL
	case FULL_SAIL:
		return ALL_PLAIN_SAIL
	case EXTRA_SAIL:
		return FULL_SAIL
	}
	panic(fmt.Sprintf("assert(SailSetting != %d)", s))
}
