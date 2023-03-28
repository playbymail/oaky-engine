// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package ship

import (
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"github.com/playbymail/oaky-engine/engine/sail"
	"github.com/playbymail/oaky-engine/engine/sailing_chart"
)

type Ship struct {
	name         string
	rate         engine.ShipRate
	sailingChart *sailing_chart.Chart
	quality      engine.ShipQuality
	coordinates  engine.Coordinates
	heading      float64 // radians, clockwise, 0 is North
	sailSetting  sail.Setting
}

// New returns an initialized ship
func New(name string, rate engine.ShipRate, options ...Option) (Ship, error) {
	s := Ship{name: name, rate: rate}
	for _, chart := range sailing_chart.Charts() {
		for _, rate := range chart.Rates {
			if s.rate == rate {
				s.sailingChart = chart
				break
			}
		}
	}
	for _, option := range options {
		if err := option(&s); err != nil {
			return Ship{}, err
		}
	}
	return s, nil
}

func (s Ship) Coordinates() engine.Coordinates {
	return s.coordinates
}

func (s Ship) Heading() float64 {
	return s.heading
}

func (s Ship) IncreaseSail() (Ship, error) {
	tmp := s.sailSetting.AddSail()
	if tmp == s.sailSetting {
		return s, fmt.Errorf("invalid sail setting")
	}
	s.sailSetting = tmp
	return s, nil
}

func (s Ship) SailingChart() *sailing_chart.Chart {
	return s.sailingChart
}

func (s Ship) SailSetting() sail.Setting {
	return s.sailSetting
}

func (s Ship) SetHeading(radians float64) (Ship, error) {
	s.heading = radians
	return s, nil
}

func (s Ship) SetSails(setting sail.Setting) (Ship, error) {
	s.sailSetting = setting
	return s, nil
}
