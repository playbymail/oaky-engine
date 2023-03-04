// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

// WindSpeed is sometimes called "WindStrength" in the rule book.

import (
	"bytes"
	"fmt"
)

const (
	// Strength of the wind affects sailing ability.
	// The stronger the wind, the faster the vessel will move; but a
	// very strong wind can destroy a ship if it sets too much sail.
	//
	// These constants define the possible Wind Strengths.

	NO_WIND WindSpeed = 0

	// Ship will not move unless it is moved under oars, or are being towed by boats.
	CALM = 1

	// Ship may "catch a wind" and move, while others will have to row, or be towed by boats.
	LIGHT_GUSTS = 2

	// Light wind.
	LIGHT_WIND = 3

	// Normal breeze.
	NORMAL_BREEZE = 4

	// Ship may not use Sail Setting EXTRA_SAIL without losing masts.
	HEAVY_BREEZE = 5

	// Ship may not use Sail Setting FULL_SAIL or higher without losing masts.
	// Lower gunports may be awash.
	GALE = 6

	// Ship may not use Sail Setting ALL_PLAIN_SAIL or higher without losing masts.
	// Lower gunports may be awash.
	STORM = 7

	// Ship may not use Sail Setting FIGHTING_SAIL or higher without losing masts.
	// No combat is allowed.
	HURRICANE = 8

	// sizeofWindSpeed is used for sizing arrays.
	// It must be one more than the largest enum value.
	sizeofWindSpeed = HURRICANE + 1
)

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
