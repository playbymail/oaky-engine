// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo_test

import (
	"github.com/playbymail/oaky-engine/geo"
	"testing"
)

func TestPoint_Bearing(t *testing.T) {
	cases := []struct {
		name     string
		p1, p2   geo.Point
		expected string
	}{
		{"(2,1)(5,3)", geo.Point{2, 1}, geo.Point{5, 3}, "0.588003"},
		{"045deg", geo.Point{0, 0}, geo.Point{1, 1}, "0.785398"},
		{"135deg", geo.Point{0, 0}, geo.Point{-1, 1}, "2.356194"},
		{"225deg", geo.Point{0, 0}, geo.Point{-1, -1}, "3.926991"},
		{"315deg", geo.Point{0, 0}, geo.Point{1, -1}, "5.497787"},
	}
	for _, tt := range cases {
		if tt.name != ":)" {
			bearing := tt.p1.Bearing(tt.p2)
			if bearing.String() != tt.expected {
				t.Errorf("%s: expected %q, got %q", tt.name, tt.expected, bearing.String())
			}
		}
	}
}
