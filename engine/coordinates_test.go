// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"math"
	"testing"
)

func TestCoordinates_Bearing(t *testing.T) {
	type testCase struct {
		a, b        engine.Coordinates
		distance    string
		degreesTo   string
		degreesFrom string
	}
	for n, tc := range []testCase{
		{a: engine.Coordinates{}, b: engine.Coordinates{3, 2}, distance: "3.6056", degreesTo: "56.3099", degreesFrom: "236.3099"},
	} {
		a, b := tc.a, tc.b
		bearing, distance := a.BearingTo(b)
		if degreesTo := fmt.Sprintf("%.4f", bearing*180.0/math.Pi); degreesTo != tc.degreesTo {
			t.Errorf("%d: %s -> %s: bearing: want %s: got %s\n", n+1, a, b, tc.degreesTo, degreesTo)
		}
		if distanceTo := fmt.Sprintf("%.4f", distance); distanceTo != tc.distance {
			t.Errorf("%d: %s -> %s: distance: want %s: got %s\n", n+1, a, b, tc.distance, distanceTo)
		}
		bearing, distance = a.BearingFrom(b)
		if degreesFrom := fmt.Sprintf("%.4f", bearing*180.0/math.Pi); degreesFrom != tc.degreesFrom {
			t.Errorf("%d: %s <- %s: bearing: want %s: got %s\n", n+1, a, b, tc.degreesFrom, degreesFrom)
		}
		if distanceFrom := fmt.Sprintf("%.4f", distance); distanceFrom != tc.distance {
			t.Errorf("%d: %s <- %s: distance: want %s: got %s\n", n+1, a, b, tc.distance, distanceFrom)
		}
	}
}
