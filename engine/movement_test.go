// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"testing"
)

type WIND struct {
	From    float64 // radians, clockwise, 0 is north
	Bearing engine.WindBearing
	Speed   engine.WindSpeed
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
	wind := WIND{
		From:  engine.DegreesToRadians(0), // from the north
		Speed: engine.LIGHT_WIND,
	}
	if wind.Speed.String() != "LIGHT_WIND" {
		t.Errorf("wind: speed: want LIGHT_WIND: got %q\n", wind.Speed.String())
	}

	// USS Constitution is a 4th rate frigate
	constitution := engine.Ship{
		Name:        "USS Constitution",
		Rate:        engine.FRIGATE_4TH_RATE,
		Quality:     engine.AVERAGE,
		Coordinates: engine.Coordinates{0, 0},
	}
	if ok := constitution.SetSailingChart(); !ok {
		t.Errorf("constitution: setSailingChart: want ok: got !ok\n")
	} else if constitution.SailingChart.Name != "SAILING_CHART_II" {
		t.Errorf("constitution: sailingChart: want SAILING_CHART_II: got %q\n", constitution.SailingChart.Name)
	}

	// Constitution stars with fighting sail and heading East
	constitution.SailSetting = engine.FIGHTING_SAIL
	if constitution.SailSetting.String() != "FIGHTING_SAIL" {
		t.Errorf("constitution: 1st leg: sail: want FIGHTING_SAIL: got %q\n", constitution.SailSetting.String())
	}
	constitution.Heading = engine.DegreesToRadians(90) // east
	if fmt.Sprintf("%.04f", constitution.Heading) != "1.5708" {
		t.Errorf("constitution: 1st leg: turn: heading: want 1.5708: got %s\n", fmt.Sprintf("%.04f", constitution.Heading))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading east, it should be a broad reach.
	wind.Bearing = engine.WB_BROAD_REACH // todo: calculate this
	if wind.Bearing.String() != "BROAD_REACH" {
		t.Errorf("constitution: 1st leg: windBearing: want BROAD_REACH: got %q\n", wind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 32
	speed := constitution.SailingChart.WindSpeed[wind.Speed].SailSetting[constitution.SailSetting].WindBearing[wind.Bearing].Speed
	if speed != 32 {
		t.Errorf("constitution: 1st leg: speed: want 32: got %d\n", speed)
	}

	// Turn the Constitution to Turns to ESE
	constitution.Heading = engine.DegreesToRadians(112) // ESE is actually 112.5
	if fmt.Sprintf("%.04f", constitution.Heading) != "1.9548" {
		t.Errorf("constitution: 1st leg: turn: heading: want 1.9548: got %s\n", fmt.Sprintf("%.04f", constitution.Heading))
	}

	// Increase sail to ALL_PLAIN_SAIL
	constitution.IncreaseSail()
	if constitution.SailSetting.String() != "ALL_PLAIN_SAIL" {
		t.Errorf("constitution: 2nd leg: sail: want ALL_PLAIN_SAIL: got %s\n", constitution.SailSetting.String())
	}

	// calculate the wind bearing. if it is from the north and the ship is heading ESE, it should be a quarter reach.
	wind.Bearing = engine.WB_QUARTER_REACH // todo: calculate this
	if wind.Bearing.String() != "QUARTER_REACH" {
		t.Errorf("constitution: 2nd leg: windBearing: want QUARTER_REACH: got %q\n", wind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 72
	speed = constitution.SailingChart.WindSpeed[wind.Speed].SailSetting[constitution.SailSetting].WindBearing[wind.Bearing].Speed
	if speed != 72 {
		t.Errorf("constitution: 2nd leg: speed: want 72: got %d\n", speed)
	}

	// turn back to east
	constitution.Heading = engine.DegreesToRadians(90) // east
	if fmt.Sprintf("%.04f", constitution.Heading) != "1.5708" {
		t.Errorf("constitution: 3rd leg: turn: heading: want 1.5708: got %s\n", fmt.Sprintf("%.04f", constitution.Heading))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading east, it should be a broad reach.
	wind.Bearing = engine.WB_BROAD_REACH // todo: calculate this
	if wind.Bearing.String() != "BROAD_REACH" {
		t.Errorf("constitution: 3rd leg: windBearing: want BROAD_REACH: got %q\n", wind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 56
	speed = constitution.SailingChart.WindSpeed[wind.Speed].SailSetting[constitution.SailSetting].WindBearing[wind.Bearing].Speed
	if speed != 56 {
		t.Errorf("constitution: 3rd leg: speed: want 56: got %d\n", speed)
	}

	constitution.Heading = engine.DegreesToRadians(45) // north-east
	if fmt.Sprintf("%.04f", constitution.Heading) != "0.7854" {
		t.Errorf("constitution: 4th leg: turn: heading: want 0.7854: got %s\n", fmt.Sprintf("%.04f", constitution.Heading))
	}

	// calculate the wind bearing. if it is from the north and the ship is heading north-east, it should be beating.
	wind.Bearing = engine.WB_BEATING // todo: calculate this
	if wind.Bearing.String() != "BEATING" {
		t.Errorf("constitution: 4th leg: windBearing: want BEATING: got %q\n", wind.Bearing.String())
	}

	// determine the speed from the sailing chart. for this ship with this wind, this reach, and these sails, we should get 28
	speed = constitution.SailingChart.WindSpeed[wind.Speed].SailSetting[constitution.SailSetting].WindBearing[wind.Bearing].Speed
	if speed != 28 {
		t.Errorf("constitution: 4th leg: speed: want 28: got %d\n", speed)
	}
}
