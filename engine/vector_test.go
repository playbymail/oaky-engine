// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"github.com/playbymail/oaky-engine/engine"
	"testing"
)

func TestVector(t *testing.T) {
	for _, tc := range []struct {
		id     int
		v1, v2 engine.Vector
		sum    engine.Vector
	}{
		{1, engine.Vector{2, 4}, engine.Vector{1, 5}, engine.Vector{3, 9}},
	} {
		sum := tc.v1.Add(tc.v2)
		if sum.String() != tc.sum.String() {
			t.Errorf("sum: %d: want %s: got %s\n", tc.id, tc.sum.String(), sum.String())
		}
	}
}
