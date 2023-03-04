// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package engine

/*

 */

type Chart struct {
	Name                 string
	WindBearingSpeedSail [sizeofWindBearing][sizeofWindSpeed][sizeofSailSetting]SailingSpeed
}

type SailingSpeed struct {
	Allowed bool // false if the entry is forbidden or undefined
	Splat   bool // true if the entry is splatted
	Speed   int
}

type SailingChart struct {
	SAILING_CHART_I   *SAILING_CHART `json:"SAILING_CHART_I"`
	SAILING_CHART_II  *SAILING_CHART `json:"SAILING_CHART_II"`
	SAILING_CHART_III *SAILING_CHART `json:"SAILING_CHART_III"`
	SAILING_CHART_IV  *SAILING_CHART `json:"SAILING_CHART_IV"`
}

type SAILING_CHART struct {
	Rates            []string      `json:"rates"`
	Notes            string        `json:"notes"`
	WB_ASTERN        *SAIL_SETTING `json:"WIND_ASTERN"`
	WB_QUARTER_REACH *SAIL_SETTING `json:"QUARTER_REACH"`
	WB_BROAD_REACH   *SAIL_SETTING `json:"BROAD_REACH"`
	WB_BEATING       *SAIL_SETTING `json:"BEATING"`
}

type SAIL_SETTING struct {
	MINIMUM_SAIL   *WIND_SPEED `json:"MINIMUM_SAIL"`
	FIGHTING_SAIL  *WIND_SPEED `json:"FIGHTING_SAIL"`
	ALL_PLAIN_SAIL *WIND_SPEED `json:"ALL_PLAIN_SAIL"`
	FULL_SAIL      *WIND_SPEED `json:"FULL_SAIL"`
	EXTRA_SAIL     *WIND_SPEED `json:"EXTRA_SAIL"`
}

type WIND_SPEED struct {
	LIGHT_WIND    []int `json:"LIGHT_WIND"`
	NORMAL_BREEZE []int `json:"NORMAL_BREEZE"`
	HEAVY_BREEZE  []int `json:"HEAVY_BREEZE"`
	GALE          []int `json:"GALE"`
	STORM         []int `json:"STORM"`
	HURRICANE     []int `json:"HURRICANE"`
}
