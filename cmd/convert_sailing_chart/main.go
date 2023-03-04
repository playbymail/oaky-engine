// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements a tool to convert the Sailing Chart from JSON to
// the engine's data structures.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/playbymail/oaky-engine/engine"
	"log"
	"os"
)

func main() {
	sailingChart, err := loadChartTables("sailing_chart_tables.json")
	if err != nil {
		log.Fatal(err)
	}

	charts := []*Chart{{Name: "Sailing Chart I"}, {Name: "Sailing Chart II"}, {Name: "Sailing Chart III"}, {Name: "Sailing Chart IV"}, {Name: "Sailing Chart V"}}
	for n, chart := range charts {
		var sc *ChartWindBearing
		switch n {
		case 0:
			sc = sailingChart.SAILING_CHART_I
		case 1:
			sc = sailingChart.SAILING_CHART_II
		case 2:
			sc = sailingChart.SAILING_CHART_III
		case 3:
			sc = sailingChart.SAILING_CHART_IV
		case 4:
			sc = sailingChart.SAILING_CHART_V
		}
		if sc != nil {
			for _, a := range []struct {
				wind_bearing engine.WindBearing
				data         *ChartSailSetting
			}{
				{engine.WB_ASTERN, sc.WB_ASTERN},
				{engine.WB_QUARTER_REACH, sc.WB_QUARTER_REACH},
				{engine.WB_BROAD_REACH, sc.WB_BROAD_REACH},
				{engine.WB_BEATING, sc.WB_BEATING},
			} {
				if a.data != nil {
					for _, b := range []struct {
						sail_setting engine.SailSetting
						data         *ChartWindSpeed
					}{
						{engine.MINIMUM_SAIL, a.data.MINIMUM_SAIL},
						{engine.FIGHTING_SAIL, a.data.FIGHTING_SAIL},
						{engine.ALL_PLAIN_SAIL, a.data.ALL_PLAIN_SAIL},
						{engine.FULL_SAIL, a.data.FULL_SAIL},
						{engine.EXTRA_SAIL, a.data.EXTRA_SAIL},
					} {
						if b.data != nil {
							for _, c := range []struct {
								wind_speed engine.WindSpeed
								data       [4]int
							}{
								{engine.LIGHT_WIND, b.data.LIGHT_WIND},
								{engine.NORMAL_BREEZE, b.data.NORMAL_BREEZE},
								{engine.HEAVY_BREEZE, b.data.HEAVY_BREEZE},
								{engine.GALE, b.data.GALE},
								{engine.STORM, b.data.STORM},
								{engine.HURRICANE, b.data.HURRICANE},
							} {
								chart.WindBearingSpeedSail[a.wind_bearing][c.wind_speed][b.sail_setting] = SailingSpeed{
									Allowed: c.data[3] != 0,
									Splat:   c.data[3] == 999,
									Speed:   c.data[1],
								}
							}
						}
					}
				}
			}
		}
	}

	bb := &bytes.Buffer{}
	_, _ = fmt.Fprintf(bb, "// Generated code; DO NOT EDIT\n\n")
	_, _ = fmt.Fprintf(bb, "package gumbo\n\n")
	_, _ = fmt.Fprintf(bb, "import \"github.com/playbymail/oaky-engine/engine\"\n\n")
	_, _ = fmt.Fprintf(bb, "const (\n")
	_, _ = fmt.Fprintf(bb, "\tSAILING_CHART_I = 0\n")
	_, _ = fmt.Fprintf(bb, "\tSAILING_CHART_II = 1\n")
	_, _ = fmt.Fprintf(bb, "\tSAILING_CHART_III = 2\n")
	_, _ = fmt.Fprintf(bb, "\tSAILING_CHART_IV = 3\n")
	_, _ = fmt.Fprintf(bb, "\tSAILING_CHART_V = 4\n")
	_, _ = fmt.Fprintf(bb, ")\n\n")
	_, _ = fmt.Fprintf(bb, "func newCharts() []engine.Chart {\n")
	_, _ = fmt.Fprintf(bb, "\tchart := make([]engine.Chart, 5, 5)\n")
	for n, chart := range charts {
		var sc string
		switch n {
		case 0:
			sc = "SAILING_CHART_I"
		case 1:
			sc = "SAILING_CHART_II"
		case 2:
			sc = "SAILING_CHART_III"
		case 3:
			sc = "SAILING_CHART_IV"
		case 4:
			sc = "SAILING_CHART_V"
		}
		_, _ = fmt.Fprintf(bb, "\n\tchart[%-17s].Name = %q\n", sc, sc)
		for wind_bearing := engine.WB_ASTERN; wind_bearing <= engine.WB_BEATING; wind_bearing++ {
			for wind_speed := engine.LIGHT_WIND; wind_speed <= engine.HURRICANE; wind_speed++ {
				for sail_setting := engine.NO_SAIL; sail_setting <= engine.EXTRA_SAIL; sail_setting++ {
					if chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Speed == 0 {
						continue
					} else if !chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Allowed {
						continue
					}
					if chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Splat {
						_, _ = fmt.Fprintf(bb, "\tchart[%-17s].WindBearingSpeedSail[engine.WB_%-14s][engine.%-14s][engine.%-14s] = engine.SailingSpeed{Speed: %3d, Allowed:true, Splat: true}\n",
							sc,
							engine.WindBearing(wind_bearing).String(),
							engine.WindSpeed(wind_speed).String(),
							engine.SailSetting(sail_setting).String(),
							chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Speed)
					} else {
						_, _ = fmt.Fprintf(bb, "\tchart[%-17s].WindBearingSpeedSail[engine.WB_%-14s][engine.%-14s][engine.%-14s] = engine.SailingSpeed{Speed: %3d, Allowed:true}\n",
							sc,
							engine.WindBearing(wind_bearing).String(),
							engine.WindSpeed(wind_speed).String(),
							engine.SailSetting(sail_setting).String(),
							chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Speed)
					}
				}
			}
		}
	}
	_, _ = fmt.Fprintf(bb, "\n\treturn chart\n")
	_, _ = fmt.Fprintf(bb, "}\n")
	if err := os.WriteFile("D:\\playbymail\\oaky\\engine\\testdata\\sailing_chart_gumbo.go", bb.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

func loadChartTables(name string) (SailingChartTables, error) {
	var sailingChart SailingChartTables
	if data, err := os.ReadFile(name); err != nil {
		return sailingChart, err
	} else if err = json.Unmarshal(data, &sailingChart); err != nil {
		return sailingChart, err
	}
	return sailingChart, nil
}
