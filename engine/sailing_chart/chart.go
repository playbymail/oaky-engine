// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package sailing_chart

import (
	"github.com/playbymail/oaky-engine/engine"
	"github.com/playbymail/oaky-engine/engine/sail"
	"github.com/playbymail/oaky-engine/engine/wind"
)

// Chart holds the sailing chart data from the rulebook.
type Chart struct {
	Name      string
	Rates     []engine.ShipRate
	WindSpeed [wind.MAX_SPEED + 1]struct {
		SailSetting [sail.MAX_SETTING + 1]struct {
			Allowed     bool
			WindBearing [wind.MAX_BEARING + 1]struct {
				Splat bool
				Speed int
			}
		}
	}
}
