// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/playbymail/oaky-engine/engine"
	"github.com/playbymail/oaky-engine/engine/wind"
)

type Chart struct {
	Name                 string
	Rates                []string
	WindBearingSpeedSail [wind.HURRICANE + 1][engine.EXTRA_SAIL + 1][wind.WB_BEATING + 1]SailingSpeed
}

type SailingSpeed struct {
	Allowed bool // false if the entry is forbidden or undefined
	Splat   bool // true if the entry is splatted
	Speed   int
}

type SailingChartTables struct {
	SAILING_CHART_I   *ChartWindBearing `json:"SAILING_CHART_I"`
	SAILING_CHART_II  *ChartWindBearing `json:"SAILING_CHART_II"`
	SAILING_CHART_III *ChartWindBearing `json:"SAILING_CHART_III"`
	SAILING_CHART_IV  *ChartWindBearing `json:"SAILING_CHART_IV"`
	SAILING_CHART_V   *ChartWindBearing `json:"SAILING_CHART_V"`
}

type ChartWindBearing struct {
	RATES            []string          `json:"RATES"`
	WB_ASTERN        *ChartSailSetting `json:"WIND_ASTERN"`
	WB_QUARTER_REACH *ChartSailSetting `json:"QUARTER_REACH"`
	WB_BROAD_REACH   *ChartSailSetting `json:"BROAD_REACH"`
	WB_BEATING       *ChartSailSetting `json:"BEATING"`
}

type ChartSailSetting struct {
	MINIMUM_SAIL   *ChartWindSpeed `json:"MINIMUM_SAIL"`
	FIGHTING_SAIL  *ChartWindSpeed `json:"FIGHTING_SAIL"`
	ALL_PLAIN_SAIL *ChartWindSpeed `json:"ALL_PLAIN_SAIL"`
	FULL_SAIL      *ChartWindSpeed `json:"FULL_SAIL"`
	EXTRA_SAIL     *ChartWindSpeed `json:"EXTRA_SAIL"`
}

type ChartWindSpeed struct {
	LIGHT_WIND    [4]int `json:"LIGHT_WIND"`
	NORMAL_BREEZE [4]int `json:"NORMAL_BREEZE"`
	HEAVY_BREEZE  [4]int `json:"HEAVY_BREEZE"`
	GALE          [4]int `json:"GALE"`
	STORM         [4]int `json:"STORM"`
	HURRICANE     [4]int `json:"HURRICANE"`
}
