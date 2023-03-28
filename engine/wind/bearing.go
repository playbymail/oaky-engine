// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package wind

import "fmt"

type Bearing int

const (
	ASTERN        Bearing = 0
	QUARTER_REACH         = 1
	BROAD_REACH           = 2
	BEATING               = 3
)

// MAX_BEARING is the largest value for Bearing.
// It is meant to be used for sizing arrays.
const MAX_BEARING = BEATING

func (b Bearing) String() string {
	switch b {
	case ASTERN:
		return "ASTERN"
	case QUARTER_REACH:
		return "QUARTER_REACH"
	case BROAD_REACH:
		return "BROAD_REACH"
	case BEATING:
		return "BEATING"
	}
	panic(fmt.Sprintf("assert(Bearing != %d)", b))
}
