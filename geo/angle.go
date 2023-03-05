// oaky-engine: a game engine for Walter Jon William's "Heart of Oak"
// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package geo

import (
	"fmt"
	"math"
)

// Angle is an immutable one-dimensional value with a unit of radians.
// It measures the offset, counterclockwise, from the positive X-axis.
// A normalized range is 0 <= Angle < 2π
type Angle float64

const twoPi = math.Pi + math.Pi

// The names, intervals, and measuring units are shown in the table below:
//
// |unit  |zero angle|acute angle|right angle|obtuse angle|straight angle|reflex angle|perigon|
// |turn  |0 turn    |0...¼ turn |¼ turn     |¼...½ turn  |½ turn        |½...1 turn  |1 turn |
// |radian|0 rad     |0...½π rad |½π rad     |½π...π rad  |π rad         |π...2π rad  |2π rad |
// |degree|0°        |0...90°    |90°        |90...180°   |180°          |180...360°  |360°   |
//

// approx returns a representation of the floating point number when
// testing for equality. the number is rounded to six decimal places
// using whatever rounding is implemented by the compiler/system.
// the approximation can't be used for ordering.
func approx(f float64) string {
	return fmt.Sprintf("%.6f", f)
}

func (a Angle) approx() string {
	return approx(float64(a))
}

func RadiansToAngle(r float64) Angle {
	return Angle(r).Normalize()
}

func (a Angle) Add(b Angle) Angle {
	return (a + b).Normalize()
}

func (a Angle) Equals(b Angle) bool {
	return a.approx() == b.approx()
}

func (a Angle) IsAcute() bool {
	return a < math.Pi/2
}

func (a Angle) IsOblique() bool {
	return a.IsAcute() || a.IsObtuse()
}

func (a Angle) IsObtuse() bool {
	return math.Pi/2 < a && a < math.Pi
}

// IsPerigon returns true if the angle is 360°.
// Wikipedia: An angle equal to 1 turn (360° or 2π radians)
// is called a full angle, complete angle, round angle or perigon.
func (a Angle) IsPerigon() bool {
	return a.approx() == approx(math.Pi*2)
}

// IsReflex returns true if the angle is between 180° and 360°.
// Wikipedia: An angle larger than a straight angle but less
// than 1 turn (between 180° and 360°) is called a reflex angle.
func (a Angle) IsReflex() bool {
	return math.Pi < a && a < math.Pi*2
}

// IsRight returns true if the angle is 90°.
// Wikipedia: An angle equal to ¼ turn (90° or π/2 radians)
// is called a right angle.
func (a Angle) IsRight() bool {
	return a.approx() == approx(math.Pi/2)
}

// IsStraight returns true if the angle is 180°.
// Wikipedia: An angle equal to ½ turn (180° or π radians)
// is called a straight angle.
func (a Angle) IsStraight() bool {
	return a.approx() == approx(math.Pi)
}

// IsZero returns true if the angle is 0°
func (a Angle) IsZero() bool {
	return a.approx() == approx(0)
}

// Normalize constrains angle to a range of 0 <= Angle < 2π.
// It does this very, very inefficiently.
func (a Angle) Normalize() Angle {
	for a < 0 {
		a = a + twoPi
	}
	for a >= math.Pi*2 {
		a = a - twoPi
	}
	return a
}

// Sub returns the difference between two angles.
func (a Angle) Sub(b Angle) Angle {
	return a - b //.Normalize()
}

// RelativeAngle translates an absolute angle to an angle relative to the vector.
func (a Angle) RelativeAngle(b Angle) Angle {
	return (twoPi - (a - b)).Normalize()
}

// String implements the Stringer interface
func (a Angle) String() string {
	return a.approx()
}

// ToDegrees converts an angle to degrees.
func (a Angle) ToDegrees() Degrees {
	degrees, minutes, seconds := 359, 59, 59
	return DMSToDegrees(degrees, minutes, seconds)
}
