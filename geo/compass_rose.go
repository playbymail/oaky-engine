// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo

import (
	"fmt"
	"math"
)

var compassRosePoints [32][3]string

func init() {
	compassRosePoints = [32][3]string{
		{"E", "East", "Dead Ahead"},
		{"EbN", "East by North", "1 point off the Port Bow"},
		{"ENE", "East NorthEast", "2 points off the Port Bow"},
		{"NEbE", "NorthEast by East", "3 points off the Port Bow"},
		{"NE", "NorthEast", "4 points off the Port Bow"},
		{"NEbN", "NorthEast by North", "5 points off the Port Bow"},
		{"NNE", "North NorthEast", "6 points off the Port Bow"},
		{"NbE", "North by East", "7 point off the Port Bow"},
		{"N", "North", "Off the Port Side"},
		{"NbW", "North by West", "7 points off the Port Stern"},
		{"NNW", "North North NorthWest", "6 points off the Port Stern"},
		{"NWbN", "NorthWest by North", "5 points off the Port Stern"},
		{"NW", "NorthWest", "4 points off the Port Stern"},
		{"NWbW", "NorthWest by West", "3 points off the Port Stern"},
		{"WNW", "West NorthWest", "2 points off the Port Stern"},
		{"WbN", "West by North", "1 point off the Port Stern"},
		{"W", "West", "Astern"},
		{"WbS", "West by South", "1 point off the Starboard Stern"},
		{"WSW", "West SouthWest", "2 points off the Starboard Stern"},
		{"SWbW", "SouthWest by West", "3 points off the Starboard Stern"},
		{"SW", "SouthWest", "4 points off the Starboard Stern"},
		{"SWbS", "SouthWest by South", "5 points off the Starboard Stern"},
		{"SSW", "South SouthWest", "6 points off the Starboard Stern"},
		{"SbW", "South by West", "7 point off the Starboard Stern"},
		{"S", "South", "Off the Starboard Side"},
		{"SbE", "South by East", "7 points off the Starboard Bow"},
		{"SSE", "South SouthEast", "6 points off the Starboard Bow"},
		{"SEbS", "SouthEast by South", "5 points off the Starboard Bow"},
		{"SE", "SouthEast", "4 points off the Starboard Bow"},
		{"SEbE", "SouthEast by East", "3 points off the Starboard Bow"},
		{"ESE", "East SouthEast", "2 points off the Starboard Bow"},
		{"EbS", "East by South", "1 point off the Starboard Bow"},
	}
}

// FromCompassRose converts the text description of the compass bearing
// to one of the 32 points on the compass.
func FromCompassRose(s string) (Angle, error) {
	switch s {
	case "E":
		return Angle(0 * math.Pi * 2 / 32), nil
	case "EbN":
		return Angle(1 * math.Pi * 2 / 32), nil
	case "ENE":
		return Angle(2 * math.Pi * 2 / 32), nil
	case "NEbE":
		return Angle(3 * math.Pi * 2 / 32), nil
	case "NE":
		return Angle(4 * math.Pi * 2 / 32), nil
	case "NEbN":
		return Angle(5 * math.Pi * 2 / 32), nil
	case "NNE":
		return Angle(6 * math.Pi * 2 / 32), nil
	case "NbE":
		return Angle(7 * math.Pi * 2 / 32), nil
	case "N":
		return Angle(8 * math.Pi * 2 / 32), nil
	case "NbW":
		return Angle(9 * math.Pi * 2 / 32), nil
	case "NNW":
		return Angle(10 * math.Pi * 2 / 32), nil
	case "NWbN":
		return Angle(11 * math.Pi * 2 / 32), nil
	case "NW":
		return Angle(12 * math.Pi * 2 / 32), nil
	case "NWbW":
		return Angle(13 * math.Pi * 2 / 32), nil
	case "WNW":
		return Angle(14 * math.Pi * 2 / 32), nil
	case "WbN":
		return Angle(15 * math.Pi * 2 / 32), nil
	case "W":
		return Angle(16 * math.Pi * 2 / 32), nil
	case "WbS":
		return Angle(17 * math.Pi * 2 / 32), nil
	case "WSW":
		return Angle(18 * math.Pi * 2 / 32), nil
	case "SWbW":
		return Angle(19 * math.Pi * 2 / 32), nil
	case "SW":
		return Angle(20 * math.Pi * 2 / 32), nil
	case "SWbS":
		return Angle(21 * math.Pi * 2 / 32), nil
	case "SSW":
		return Angle(22 * math.Pi * 2 / 32), nil
	case "SbW":
		return Angle(23 * math.Pi * 2 / 32), nil
	case "S":
		return Angle(24 * math.Pi * 2 / 32), nil
	case "SbE":
		return Angle(25 * math.Pi * 2 / 32), nil
	case "SSE":
		return Angle(26 * math.Pi * 2 / 32), nil
	case "SEbS":
		return Angle(27 * math.Pi * 2 / 32), nil
	case "SE":
		return Angle(28 * math.Pi * 2 / 32), nil
	case "SEbE":
		return Angle(29 * math.Pi * 2 / 32), nil
	case "ESE":
		return Angle(30 * math.Pi * 2 / 32), nil
	case "EbS":
		return Angle(31 * math.Pi * 2 / 32), nil
	}
	return 0, fmt.Errorf("invalid compass rose value")
}
