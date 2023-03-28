// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"github.com/playbymail/oaky-engine/engine/sail"
	"github.com/playbymail/oaky-engine/engine/ship"
	"github.com/playbymail/oaky-engine/engine/wind"
	"testing"
)

type WIND struct {
	From    float64 // radians, clockwise, 0 is north
	Bearing wind.Bearing
	Speed   wind.Speed
}

func TestMovement_ConstitutionExample(t *testing.T) {
	// Constitution Example is taken from the rule book.
	//   The wind is from N, and the CONSTITUTION, a 4th Rate ship, is
	//   heading E. This is a broad reach. CONSTITUTION is at
	//   Sail Condition 2 (Fighting Sail), at a Wind Strength of 3.
	//   The speed is 32 meters / minute for the 1st leg.
	//
	//   CONSTITUTION then increases its Sail Setting to 3 while turning
	//   to ESE, It is now at Quarter Reach with a new sail setting, and
	//   speed is now 72 meters / minute.
	//
	//   It then turns to point E once more, and will begin the next turn
	//   in a Broad Reach.
	//
	//   Wind is from N at Strength 3, CONSTITUTION is Broad Reaching at
	//   Sail Setting 3. During its 3rd leg, CONSTITUTION speed is
	//   56 meters / minute. She alters course to NE. It is now Beating,
	//   and moves at 28 meters / minutes for the 4th leg.

	// light wind from the north
	theWind := WIND{
		From:  engine.DegreesToRadians(0), // from the north
		Speed: wind.LIGHT_WIND,
	}
	if theWind.Speed.String() != "LIGHT_WIND" {
		t.Errorf("wind: speed: want LIGHT_WIND: got %q\n", theWind.Speed.String())
	}

	// USS Constitution is a 4th rate frigate
	opts := []ship.Option{
		ship.OptQuality(engine.AVERAGE),
		ship.OptCoordinates(engine.Coordinates{0, 0}),
	}
	constitution, err := ship.New("USS Constitution", engine.FRIGATE_4TH_RATE, opts...)
	if err != nil {
		t.Fatalf("constitution: new: error: want nil: got %v\n", err)
	} else if constitution.SailingChart().Name != "CHART_II" {
		t.Errorf("constitution: sailingChart: want CHART_II: got %q\n", constitution.SailingChart().Name)
	}

	// Constitution stars with fighting sail and heading East
	if constitution, err = constitution.SetSails(sail.FIGHTING_SAIL); err != nil {
		t.Errorf("constitution: 1st leg: sail: error %v\n", err)
	} else if constitution.SailSetting().String() != "FIGHTING_SAIL" {
		t.Errorf("constitution: 1st leg: sail: want FIGHTING_SAIL: got %q\n", constitution.SailSetting().String())
	}
	// east
	if constitution, err = constitution.SetHeading(engine.DegreesToRadians(90)); err != nil {
		t.Errorf("constitution: 1st leg: heading: error %v\n", err)
	} else if fmt.Sprintf("%.04f", constitution.Heading()) != "1.5708" {
		t.Errorf("constitution: 1st leg: turn: heading: want 1.5708: got %s\n", fmt.Sprintf("%.04f", constitution.Heading()))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading east, it should be a broad reach.
	theWind.Bearing = wind.BROAD_REACH // todo: calculate this
	if theWind.Bearing.String() != "BROAD_REACH" {
		t.Errorf("constitution: 1st leg: windBearing: want BROAD_REACH: got %q\n", theWind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 32
	speed := constitution.SailingChart().WindSpeed[theWind.Speed].SailSetting[constitution.SailSetting()].WindBearing[theWind.Bearing].Speed
	if speed != 32 {
		t.Errorf("constitution: 1st leg: speed: want 32: got %d\n", speed)
	}

	// Turn the Constitution to Turns to ESE // ESE is actually 112.5
	if constitution, err = constitution.SetHeading(engine.DegreesToRadians(112)); err != nil {
		t.Errorf("constitution: 1st leg: heading: error %v\n", err)
	} else if fmt.Sprintf("%.04f", constitution.Heading()) != "1.9548" {
		t.Errorf("constitution: 1st leg: turn: heading: want 1.9548: got %s\n", fmt.Sprintf("%.04f", constitution.Heading()))
	}

	// Increase sail to ALL_PLAIN_SAIL
	if constitution, err = constitution.IncreaseSail(); err != nil {
		t.Errorf("constitution: 2nd leg: sail: error %v\n", err)
	} else if constitution.SailSetting().String() != "ALL_PLAIN_SAIL" {
		t.Errorf("constitution: 2nd leg: sail: want ALL_PLAIN_SAIL: got %s\n", constitution.SailSetting().String())
	}

	// calculate the wind bearing. if it is from the north and the ship is heading ESE, it should be a quarter reach.
	theWind.Bearing = wind.QUARTER_REACH // todo: calculate this
	if theWind.Bearing.String() != "QUARTER_REACH" {
		t.Errorf("constitution: 2nd leg: windBearing: want QUARTER_REACH: got %q\n", theWind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 72
	speed = constitution.SailingChart().WindSpeed[theWind.Speed].SailSetting[constitution.SailSetting()].WindBearing[theWind.Bearing].Speed
	if speed != 72 {
		t.Errorf("constitution: 2nd leg: speed: want 72: got %d\n", speed)
	}

	// turn back to east
	if constitution, err = constitution.SetHeading(engine.DegreesToRadians(90)); err != nil {
		t.Errorf("constitution: 3rd leg: heading: error %v\n", err)
	} else if fmt.Sprintf("%.04f", constitution.Heading()) != "1.5708" {
		t.Errorf("constitution: 3rd leg: turn: heading: want 1.5708: got %s\n", fmt.Sprintf("%.04f", constitution.Heading()))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading east, it should be a broad reach.
	theWind.Bearing = wind.BROAD_REACH // todo: calculate this
	if theWind.Bearing.String() != "BROAD_REACH" {
		t.Errorf("constitution: 3rd leg: windBearing: want BROAD_REACH: got %q\n", theWind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 56
	speed = constitution.SailingChart().WindSpeed[theWind.Speed].SailSetting[constitution.SailSetting()].WindBearing[theWind.Bearing].Speed
	if speed != 56 {
		t.Errorf("constitution: 3rd leg: speed: want 56: got %d\n", speed)
	}

	// north-east
	if constitution, err = constitution.SetHeading(engine.DegreesToRadians(45)); err != nil {
		t.Errorf("constitution: 4th leg: heading: error %v\n", err)
	} else if fmt.Sprintf("%.04f", constitution.Heading()) != "0.7854" {
		t.Errorf("constitution: 4th leg: turn: heading: want 0.7854: got %s\n", fmt.Sprintf("%.04f", constitution.Heading()))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading north-east, it should be beating.
	theWind.Bearing = wind.BEATING // todo: calculate this
	if theWind.Bearing.String() != "BEATING" {
		t.Errorf("constitution: 4th leg: windBearing: want BEATING: got %q\n", theWind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 28
	speed = constitution.SailingChart().WindSpeed[theWind.Speed].SailSetting[constitution.SailSetting()].WindBearing[theWind.Bearing].Speed
	if speed != 28 {
		t.Errorf("constitution: 4th leg: speed: want 28: got %d\n", speed)
	}
}
