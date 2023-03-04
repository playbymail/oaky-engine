// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

type Ship struct {
	Name         string
	Rate         ShipRate
	SailingChart *SailingChart
	Quality      ShipQuality
	Coordinates  Coordinates
	Heading      float64 // radians, clockwise, 0 is North
	SailSetting  SailSetting
}

var sailingCharts []*SailingChart

func (s *Ship) IncreaseSail() bool {
	if tmp := s.SailSetting.AddSail(); tmp != s.SailSetting {
		s.SailSetting = tmp
		return true
	}
	return false
}

func (s *Ship) SetSailingChart() bool {
	if sailingCharts == nil {
		allSailingCharts := newSailingCharts()
		for n := range allSailingCharts {
			sailingCharts = append(sailingCharts, &allSailingCharts[n])
		}
	}

	for _, chart := range sailingCharts {
		for _, rate := range chart.Rates {
			if rate == s.Rate {
				s.SailingChart = chart
				return true
			}
		}
	}
	return false
}
