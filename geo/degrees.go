// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo

import "fmt"

type Degrees string

func (d Degrees) String() string {
	return string(d)
}

// this should be replaced with AnglesToDegrees
func DMSToDegrees(degrees, minutes, seconds int) Degrees {
	return Degrees(fmt.Sprintf("%03d° %02d′ %02d″", degrees, minutes, seconds))
}
