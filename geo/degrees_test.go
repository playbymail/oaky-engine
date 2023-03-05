// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo_test

import (
	"testing"

	"github.com/playbymail/oaky-engine/geo"
)

func Test_Degrees(t *testing.T) {
	for _, tt := range []struct {
		expect  string
		degrees int
		minutes int
		seconds int
	}{
		{"000° 00′ 00″", 0, 0, 0},
		{"000° 00′ 00″", 0, 0, 0},
		{"001° 02′ 03″", 1, 2, 3},
		{"045° 05′ 05″", 45, 5, 5},
		{"123° 45′ 56″", 123, 45, 56},
		{"180° 30′ 30″", 180, 30, 30},
		{"359° 00′ 00″", 359, 0, 0},
		{"359° 59′ 00″", 359, 59, 0},
		{"359° 59′ 59″", 359, 59, 59},
	} {
		degrees := geo.DMSToDegrees(tt.degrees, tt.minutes, tt.seconds)
		if degrees.String() != tt.expect {
			t.Errorf("expected %q, got %q", tt.expect, degrees)
		}
	}
}
