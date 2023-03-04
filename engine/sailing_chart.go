// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

type SailingChart struct {
	Name      string
	Rates     []ShipRate
	WindSpeed [sizeofWindSpeed]struct {
		SailSetting [sizeofSailSetting]struct {
			Allowed     bool
			WindBearing [sizeofWindBearing]struct {
				Splat bool
				Speed int
			}
		}
	}
}
