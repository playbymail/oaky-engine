// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

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
	var sailingChart engine.SailingChart
	if data, err := os.ReadFile("sailing_chart.json"); err != nil {
		log.Fatal(err)
	} else if err = json.Unmarshal(data, &sailingChart); err != nil {
		log.Fatal(err)
	}

	charts := []*engine.Chart{{Name: "Sailing Chart I"}, {Name: "Sailing Chart II"}, {Name: "Sailing Chart III"}, {Name: "Sailing Chart IV"}}
	for n, chart := range charts {
		if chart != nil {
			var sc *engine.SAILING_CHART
			switch n {
			case 0:
				sc = sailingChart.SAILING_CHART_I
			case 1:
				sc = sailingChart.SAILING_CHART_II
			case 2:
				sc = sailingChart.SAILING_CHART_III
			case 3:
				sc = sailingChart.SAILING_CHART_IV
			}
			if sc != nil {
				for _, a := range []struct {
					wind_bearing engine.WindBearing
					data         *engine.SAIL_SETTING
				}{
					{engine.WB_ASTERN, sc.WB_ASTERN},
					{engine.WB_QUARTER_REACH, sc.WB_QUARTER_REACH},
					{engine.WB_BROAD_REACH, sc.WB_BROAD_REACH},
					{engine.WB_BEATING, sc.WB_BEATING},
				} {
					if a.data != nil {
						for _, b := range []struct {
							sail_setting engine.SailSetting
							data         *engine.WIND_SPEED
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
									data       []int
								}{
									{engine.LIGHT_WIND, b.data.LIGHT_WIND},
									{engine.NORMAL_BREEZE, b.data.NORMAL_BREEZE},
									{engine.HEAVY_BREEZE, b.data.HEAVY_BREEZE},
									{engine.GALE, b.data.GALE},
									{engine.STORM, b.data.STORM},
									{engine.HURRICANE, b.data.HURRICANE},
								} {
									var sp engine.SailingSpeed
									switch len(c.data) {
									case 3:
										sp.Allowed, sp.Speed = true, c.data[1]
									case 4:
										sp.Allowed, sp.Speed, sp.Splat = true, c.data[1], true
									}
									//log.Printf("%-13s: allowed %-5v: %3d/%3d/%3d %v\n", c.wind_speed, sp.Allowed, 0, sp.Speed, 0, sp.Splat)
									chart.WindBearingSpeedSail[a.wind_bearing][c.wind_speed][b.sail_setting] = sp
								}
							}
						}
					}
				}
			}
		}
	}

	bb := &bytes.Buffer{}
	_, _ = fmt.Fprintf(bb, "package gumbo\n\n")
	_, _ = fmt.Fprintf(bb, "import \"github.com/playbymail/oaky-engine/engine\"\n\n")
	_, _ = fmt.Fprintf(bb, "func newCharts() []engine.Chart {\n")
	_, _ = fmt.Fprintf(bb, "\tchart := make([]engine.Chart, 5, 5)\n")
	for n, chart := range charts {
		_, _ = fmt.Fprintf(bb, "\n\tchart[%d].Name = %q\n", n+1, chart.Name)
		for wind_bearing := engine.WB_ASTERN; wind_bearing <= engine.WB_BEATING; wind_bearing++ {
			for wind_speed := engine.LIGHT_WIND; wind_speed <= engine.HURRICANE; wind_speed++ {
				for sail_setting := engine.NO_SAIL; sail_setting <= engine.EXTRA_SAIL; sail_setting++ {
					if chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Speed == 0 {
						continue
					} else if !chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Allowed {
						continue
					}
					if chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Splat {
						_, _ = fmt.Fprintf(bb, "\tchart[%d].WindBearingSpeedSail[engine.WB_%-14s][engine.%-14s][engine.%-14s] = engine.SailingSpeed{Speed: %3d, Allowed:true, Splat: true}\n",
							n+1,
							engine.WindBearing(wind_bearing).String(),
							engine.WindSpeed(wind_speed).String(),
							engine.SailSetting(sail_setting).String(),
							chart.WindBearingSpeedSail[wind_bearing][wind_speed][sail_setting].Speed)
					} else {
						_, _ = fmt.Fprintf(bb, "\tchart[%d].WindBearingSpeedSail[engine.WB_%-14s][engine.%-14s][engine.%-14s] = engine.SailingSpeed{Speed: %3d, Allowed:true}\n",
							n+1,
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
	if err := os.WriteFile("sailing_chart_gumbo.go", bb.Bytes(), 0644); err != nil {
		log.Fatal(err)
	} else if data, err := json.MarshalIndent(charts, "", "  "); err != nil {
		log.Fatal(err)
	} else if err = os.WriteFile("sailing_chart_gumbo.json", data, 0644); err != nil {
		log.Fatal(err)
	}
}
