# geo

The `geo` package implements functions to support finding location, bearing, and distance on a two-dimensional Cartesian [plane](https://en.wikipedia.org/wiki/Two-dimensional_space).

When converting to degrees, the package assumes that "north" is represented by increasing Y values and "east" by increasing X values.

## Bearings and Relative Bearings
Bearings are always measured counter-clockwise from the X-axis.

Relative bearings change the reference point from the X-axis to the vector representing heading.

## A note on comparisons

The package uses Go's `float64` for values.
When comparing for equality, we consider only the first six decimal points.

Since we're lazy, we let Go determine how to round from seven.

# Types

## Plane
All points are on the plane, which is flat.

Coordinates are given as (X,Y).
X increases going "right."
Y increases going "up."

"North" is on the positive Y-axis.
"East" is on the positive X-axis.

## Point
A [point](https://en.wikipedia.org/wiki/Point_(geometry)) is a coordinate on the plane and is represented as (X,Y).

## Angle

A one-dimensional angle, measured in radians.

## Vector

A [Euclidean vector](https://en.wikipedia.org/wiki/Euclidean_vector) consisting of length and direction.

## Degrees

Degrees represent an angle mapped to the compass rose.
The value is for display only and should never be used in calculations.

The format is "DDD° MM′ SS″", where DDD is the three-digit degrees, MM is two-digit minutes, and SS is two-digit seconds.
All values are left-padded with zeroes.

One degree is π/180 radians and has the range of 000° to 359°.

A minute is 1/60th of a degree (π/10800 radians) and has a range of 00′ to 59′.

A second is 1/60th of a minute, or 1/3600th of a degree (π/648000 radians), and has a range of 00″ to 59″.

The minimum value for a degree is 000° 00′ 00″ and the maximum is 359° 59′ 59″.

The package supports no operations on degrees other than converting angles to degrees.

## Compass Rose
The [Compass Rose](compass-rose-32-pt.svg.png) is used for reference.