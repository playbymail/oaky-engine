// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"math"
	"testing"
)

func TestBearing(t *testing.T) {
	lighthouse := engine.Coordinates{X: 0, Y: 0}
	ship := engine.Coordinates{X: -14, Y: 20}
	buoy := engine.Coordinates{X: 3, Y: 2}

	type testCase struct {
		id          int
		a, b        engine.Coordinates
		degreesTo   string
		degreesFrom string
	}

	for _, tc := range []testCase{
		{id: 1, a: lighthouse, b: buoy, degreesTo: "56.3099", degreesFrom: "236.3099"},
		{id: 2, a: buoy, b: lighthouse, degreesTo: "236.3099", degreesFrom: "56.3099"},
		{id: 3, a: lighthouse, b: ship, degreesTo: "325.0080", degreesFrom: "145.0080"},
		{id: 4, a: ship, b: lighthouse, degreesTo: "145.0080", degreesFrom: "325.0080"},
	} {
		bearing := engine.AbsoluteBearing(tc.a, tc.b)
		if degreesTo := fmt.Sprintf("%.4f", bearing*180.0/math.Pi); degreesTo != tc.degreesTo {
			t.Errorf("%d: %s -> %s: bearing: want %s: got %s\n", tc.id, tc.a, tc.b, tc.degreesTo, degreesTo)
		}
		bearing = engine.AbsoluteBearing(tc.b, tc.a)
		if degreesFrom := fmt.Sprintf("%.4f", bearing*180.0/math.Pi); degreesFrom != tc.degreesFrom {
			t.Errorf("%d: %s <- %s: bearing: want %s: got %s\n", tc.id, tc.a, tc.b, tc.degreesFrom, degreesFrom)
		}
	}
}
