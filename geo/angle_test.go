// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo_test

import (
	"fmt"
	"github.com/playbymail/oaky-engine/geo"
	"math"
	"testing"
)

func Test_Angles(t *testing.T) {
	for _, tt := range []struct {
		name    string
		radians float64
		expect  string
	}{
		{"+0π/4", 0 * math.Pi / 4, fmt.Sprintf("%.6f", 0*math.Pi/4)},
		{"+1π/4", 1 * math.Pi / 4, fmt.Sprintf("%.6f", 1*math.Pi/4)},
		{"+2π/4", 2 * math.Pi / 4, fmt.Sprintf("%.6f", 2*math.Pi/4)},
		{"+3π/4", 3 * math.Pi / 4, fmt.Sprintf("%.6f", 3*math.Pi/4)},
		{"+4π/4", 4 * math.Pi / 4, fmt.Sprintf("%.6f", 4*math.Pi/4)},
		{"+5π/4", 5 * math.Pi / 4, fmt.Sprintf("%.6f", 5*math.Pi/4)},
		{"+6π/4", 6 * math.Pi / 4, fmt.Sprintf("%.6f", 6*math.Pi/4)},
		{"+7π/4", 7 * math.Pi / 4, fmt.Sprintf("%.6f", 7*math.Pi/4)},
		{"+8π/4", 8 * math.Pi / 4, fmt.Sprintf("%.6f", 0*math.Pi/4)},
		{"+9π/4", 9 * math.Pi / 4, fmt.Sprintf("%.6f", 1*math.Pi/4)},
		{"-1π/4", -1 * math.Pi / 4, fmt.Sprintf("%.6f", 7*math.Pi/4)},
		{"-2π/4", -2 * math.Pi / 4, fmt.Sprintf("%.6f", 6*math.Pi/4)},
		{"-3π/4", -3 * math.Pi / 4, fmt.Sprintf("%.6f", 5*math.Pi/4)},
		{"-4π/4", -4 * math.Pi / 4, fmt.Sprintf("%.6f", 4*math.Pi/4)},
		{"-5π/4", -5 * math.Pi / 4, fmt.Sprintf("%.6f", 3*math.Pi/4)},
		{"-6π/4", -6 * math.Pi / 4, fmt.Sprintf("%.6f", 2*math.Pi/4)},
		{"-7π/4", -7 * math.Pi / 4, fmt.Sprintf("%.6f", 1*math.Pi/4)},
		{"-8π/4", -8 * math.Pi / 4, fmt.Sprintf("%.6f", 0*math.Pi/4)},
		{"-9π/4", -9 * math.Pi / 4, fmt.Sprintf("%.6f", 7*math.Pi/4)},
	} {
		a := geo.RadiansToAngle(tt.radians)
		if a.String() != tt.expect {
			t.Errorf("%s: expected %q: got %q", tt.name, tt.expect, a.String())
		}
	}
}
