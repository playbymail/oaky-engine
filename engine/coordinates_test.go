// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"testing"
)

func TestCoordinates(t *testing.T) {
	lighthouse := engine.Coordinates{X: 0, Y: 0}
	ship := engine.Coordinates{X: -14, Y: 20}
	buoy := engine.Coordinates{X: 3, Y: 2}
	raft := engine.Coordinates{X: 3, Y: 4}

	// verify that coordinates are immutable
	raft.Add(engine.Vector{X: 13, Y: 24})
	if raft.String() != "(3 4)" {
		t.Errorf("immutable: want (3 4): got %q\n", raft.String())
	}

	// verify distance
	for _, tc := range []struct {
		id       int
		a, b     engine.Coordinates
		distance string
	}{
		{id: 1, a: lighthouse, b: ship, distance: "24.4131"},
		{id: 2, a: ship, b: buoy, distance: "24.7588"},
		{id: 3, a: buoy, b: lighthouse, distance: "3.6056"},
		{id: 4, a: lighthouse, b: raft, distance: "5.0000"},
	} {
		if distance := fmt.Sprintf("%.04f", engine.DistanceBetween(tc.a, tc.b)); distance != tc.distance {
			t.Errorf("%d: %s -> %s: distance: want %s: got %s\n", tc.id, tc.a, tc.b, tc.distance, distance)
		}
	}
}
