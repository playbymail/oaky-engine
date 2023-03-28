// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package ship

import "github.com/playbymail/oaky-engine/engine"

type Option func(s *Ship) error

func OptCoordinates(coords engine.Coordinates) Option {
	return func(s *Ship) error {
		s.coordinates = coords
		return nil
	}
}
func OptQuality(quality engine.ShipQuality) Option {
	return func(s *Ship) error {
		s.quality = quality
		return nil
	}
}
