// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"testing"
)

func TestVector_Add(t *testing.T) {
	for _, tc := range []struct {
		id     int
		v1, v2 engine.Vector
		sum    engine.Vector
	}{
		{1, engine.Vector{2, 4}, engine.Vector{1, 5}, engine.Vector{3, 9}},
	} {
		sum := tc.v1.Add(tc.v2)
		if sum.String() != tc.sum.String() {
			t.Errorf("add: %d: want %s: got %s\n", tc.id, tc.sum.String(), sum.String())
		}
	}
}

func TestVector_Magnitude(t *testing.T) {
	for _, tc := range []struct {
		id        int
		v         engine.Vector
		magnitude float64
	}{
		{1, engine.Vector{3, 4}, 5.0},
	} {
		magnitude := tc.v.Magnitude()
		if fmt.Sprintf("%.4f", magnitude) != fmt.Sprintf("%.4f", tc.magnitude) {
			t.Errorf("magnitude: %d: want %.4f: got %.4f\n", tc.id, tc.magnitude, magnitude)
		}
	}
}

func TestVector_Normalize(t *testing.T) {
	for _, tc := range []struct {
		id            int
		v, normalized engine.Vector
	}{
		{1, engine.Vector{3, 4}, engine.Vector{0.6, 0.8}},
	} {
		normalized := tc.v.Normalize()
		if normalized.String() != tc.normalized.String() {
			t.Errorf("normalize: %d: want %s: got %s\n", tc.id, tc.normalized, normalized)
		}
	}
}

func TestVector_Scale(t *testing.T) {
	for _, tc := range []struct {
		id     int
		v      engine.Vector
		scalar float64
		scaled engine.Vector
	}{
		{1, engine.Vector{3, 4}, 1, engine.Vector{3, 4}},
		{2, engine.Vector{3, 4}, 0.2, engine.Vector{0.6, 0.8}},
		{3, engine.Vector{3, 4}, 0, engine.Vector{0, 0}},
		{1, engine.Vector{3, 4}, 2, engine.Vector{6, 8}},
	} {
		scaled := tc.v.Scale(tc.scalar)
		if scaled.String() != tc.scaled.String() {
			t.Errorf("scale: %d: want %s: got %s\n", tc.id, tc.scaled, scaled)
		}
	}
}
