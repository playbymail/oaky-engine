// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo_test

import (
	"github.com/playbymail/oaky-engine/geo"
	"math"
	"testing"
)

func TestVector_Add(t *testing.T) {
	v1 := geo.VectorFromPoint(0, 1)
	v2 := v1.Add(v1)
	if v2.X != v1.X*2 {
		t.Errorf("add: x: expected %.6f, got %.6f\n", v1.X*2, v2.X)
	}
	if v2.Y != v1.Y*2 {
		t.Errorf("add: y: expected %.6f, got %.6f\n", v1.Y*2, v2.Y)
	}
	if v2.Length() != v1.Length()*2 {
		t.Errorf("add: l: expected %.6f, got %.6f\n", v1.Length()*2, v2.Length())
	}
	if v2.Angle() != v1.Angle() {
		t.Errorf("add: a: expected %q, got %q\n", v1.Angle().String(), v2.Angle().String())
	}
}

func TestVector_RelativeAngle(t *testing.T) {
	cases := []struct {
		name     string
		v1, v2   geo.Vector
		expected geo.Angle
	}{
		{"0π     0π ", geo.Vector{}, geo.Vector{}, geo.Angle(0)},
		{" π/4   0π ", geo.Vector{A: math.Pi / 4}, geo.Vector{}, geo.Angle(7 * math.Pi / 4)},
		{" π/2   0π ", geo.Vector{A: math.Pi / 2}, geo.Vector{}, geo.Angle(3 * math.Pi / 2)},
		{" π     0π ", geo.Vector{A: math.Pi}, geo.Vector{}, geo.Angle(math.Pi)},
		{"3π/2   0π ", geo.Vector{A: 3 * math.Pi / 2}, geo.Vector{}, geo.Angle(math.Pi / 2)},
		{"0π      π/2", geo.Vector{}, geo.Vector{A: math.Pi / 2}, geo.Angle(math.Pi / 2)},
		{" π/4    π/2", geo.Vector{A: math.Pi / 4}, geo.Vector{A: math.Pi / 2}, geo.Angle(1 * math.Pi / 4)},
		{" π/2    π/2", geo.Vector{A: math.Pi / 2}, geo.Vector{A: math.Pi / 2}, geo.Angle(0)},
		{" π      π/2", geo.Vector{A: math.Pi}, geo.Vector{A: math.Pi / 2}, geo.Angle(3 * math.Pi / 2)},
		{"3π/2    π/2", geo.Vector{A: 3 * math.Pi / 2}, geo.Vector{A: math.Pi / 2}, geo.Angle(math.Pi)},
		{"3π/2  11π/8", geo.Vector{A: 3 * math.Pi / 2}, geo.Vector{A: 11 * math.Pi / 8}, geo.Angle(15 * math.Pi / 8)},
	}
	for _, tt := range cases {
		av := tt.v1.RelativeAngle(tt.v2.A)
		if !av.Equals(tt.expected) {
			t.Errorf("%s: expected %s, got %s", tt.name, tt.expected.String(), av.String())
		}
	}
}
