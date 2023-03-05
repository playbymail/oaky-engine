// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"testing"
)

func TestShipBearing(t *testing.T) {
	java := engine.Ship{
		Name:        "Java",
		Coordinates: engine.Coordinates{X: 3, Y: 3},
	}
	lighthouse := engine.Coordinates{X: 9, Y: 9}
	buoy := engine.Coordinates{X: 9, Y: 3}
	raft := engine.Coordinates{X: 9, Y: -3}

	// determine bearing from ship to other objects
	heading := 0
	java.Heading = engine.DegreesToRadians(heading)
	for _, tc := range []struct {
		id      int
		name    string
		object  engine.Coordinates
		bearing string
	}{
		{id: 1, name: "lighthouse", object: lighthouse, bearing: "0.7854"},
		{id: 2, name: "buoy", object: buoy, bearing: "1.5708"},
		{id: 3, name: "raft", object: raft, bearing: "2.3562"},
	} {
		bearing := java.BearingTo(tc.object)
		t.Errorf("%s heading %3d° -> %-8s bearing %3d° distance %.04f\n", java.Coordinates, heading, tc.object, engine.RadiansToDegrees(bearing), engine.DistanceBetween(java.Coordinates, tc.object))
		if tc.bearing != fmt.Sprintf("%.04f", bearing) {
			t.Errorf("%d: %s %3d° -> %-8s want %s got %.04f %3d°\n",
				tc.id, java.Coordinates, heading, tc.object, tc.bearing, bearing, engine.RadiansToDegrees(bearing))
		}
	}

	heading = 45
	java.Heading = engine.DegreesToRadians(heading)
	for _, tc := range []struct {
		id      int
		name    string
		object  engine.Coordinates
		bearing string
	}{
		{id: 1, name: "lighthouse", object: lighthouse, bearing: "0.0000"},
		{id: 2, name: "buoy", object: buoy, bearing: "0.7854"},
		{id: 3, name: "raft", object: raft, bearing: "1.5708"},
	} {
		bearing := java.BearingTo(tc.object)
		t.Errorf("%s heading %3d° -> %-8s bearing %3d° distance %.04f\n", java.Coordinates, heading, tc.object, engine.RadiansToDegrees(bearing), engine.DistanceBetween(java.Coordinates, tc.object))
		if tc.bearing != fmt.Sprintf("%.04f", bearing) {
			t.Errorf("%d: %s %3d° -> %-8s want %s got %.04f %3d°\n",
				tc.id, java.Coordinates, heading, tc.object, tc.bearing, bearing, engine.RadiansToDegrees(bearing))
		}
	}

	heading = 90
	java.Heading = engine.DegreesToRadians(heading)
	for _, tc := range []struct {
		id      int
		name    string
		object  engine.Coordinates
		bearing string
	}{
		{id: 1, name: "lighthouse", object: lighthouse, bearing: "5.4978"},
		{id: 2, name: "buoy", object: buoy, bearing: "0.0000"},
		{id: 3, name: "raft", object: raft, bearing: "0.7854"},
	} {
		bearing := java.BearingTo(tc.object)
		t.Errorf("%s heading %3d° -> %-8s bearing %3d° distance %.04f\n", java.Coordinates, heading, tc.object, engine.RadiansToDegrees(bearing), engine.DistanceBetween(java.Coordinates, tc.object))
		if tc.bearing != fmt.Sprintf("%.04f", bearing) {
			t.Errorf("%d: %s %3d° -> %-8s want %s got %.04f %3d°\n",
				tc.id, java.Coordinates, heading, tc.object, tc.bearing, bearing, engine.RadiansToDegrees(bearing))
		}
	}

	heading = 135
	java.Heading = engine.DegreesToRadians(heading)
	for _, tc := range []struct {
		id      int
		name    string
		object  engine.Coordinates
		bearing string
	}{
		{id: 1, name: "lighthouse", object: lighthouse, bearing: "4.7124"},
		{id: 2, name: "buoy", object: buoy, bearing: "5.4978"},
		{id: 3, name: "raft", object: raft, bearing: "0.0000"},
	} {
		bearing := java.BearingTo(tc.object)
		t.Errorf("%s heading %3d° -> %-8s bearing %3d° distance %.04f\n", java.Coordinates, heading, tc.object, engine.RadiansToDegrees(bearing), engine.DistanceBetween(java.Coordinates, tc.object))
		if tc.bearing != fmt.Sprintf("%.04f", bearing) {
			t.Errorf("%d: %s %3d° -> %-8s want %s got %.04f %3d°\n",
				tc.id, java.Coordinates, heading, tc.object, tc.bearing, bearing, engine.RadiansToDegrees(bearing))
		}
	}
}
