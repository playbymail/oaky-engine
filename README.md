# Oaky Engine: a player aid for Heart of Oak: Age of Fighting Sail

I am creating this as a play aid for Jon Williams' game, Heart of Oak.

## Copyright

The rules are Copyright &copy; 1978, 1983 by Jon Williams.
They are published by Fantasy Games Unlimited, Inc.

If you are interested in obtaining the rules, Fantasy Games Unlimited, Inc. sells a PDF.

* [Heart of Oak](https://www.fantasygamesunlimited.net/product/privateers-and-gentlemen-heart-of-oak/) $10
* [Privateers and Gentlemen](https://www.fantasygamesunlimited.net/product/privateers-and-gentlemen-rpg-rules-pdf-download/) $10

The Privateers and Gentlemen game includes a copy of the Heart of Oak rulebook.
The quality of the scan seems to be a bit better.

If FGU's site isn't available, try these guys:

* [Heart of Oak](https://www.drivethrurpg.com/product/736/Age-of-Fighting-Sail-Heart-of-Oak) $5
* [Privateers and Gentlemen](https://www.drivethrurpg.com/product/82553/Privateers-and-Gentlemen) $10

### Images
The [Compass_Card__B+W.svg](https://commons.wikimedia.org/wiki/File:Compass_Card_B+W.svg) is [CC BY-SA 3.0](http://creativecommons.org/licenses/by-sa/3.0/)

# Notes

## Ships vs Vessels
I'm going to call almost anything that floats a ship, even though the rules treat vessels differently than ships.

## Game Scale
The rules are for 1/1000 miniatures, so 1 millimeter on the table respresents 1 meter in the "real" world.

A turn represents 1 minute of time.

## Positioning
I am using a [Cartesian coordinate system](https://en.wikipedia.org/wiki/Cartesian_coordinate_system) to place objects on the map.
I'm going along with the scale in the rules, so each X and Y increment represents 1 millimeter.

The origin won't matter until we start mapping in terrain or depth.

## Heading
Heading is something like the true direction that a ship is pointing.
Again, I'm using a flat-Earth and ignoring the curvature of the Earth.

The rules represent heading using the Compass Rose.

This program uses

## Distances
Measurements are from the main mast of the ship.
I am using 64-bit floating point numbers for coordinates.
According to [Stack Overflow](https://stackoverflow.com/questions/1848700/biggest-integer-that-can-be-stored-in-a-double), that means we should be able to track objects 9,007,199,254,740,992 millimeters apart.

## Bearing
Internally, we're using radians, but we convert to the Compass Rose when reporting bearing.

I think that's kind of nice, because it hides the truth from the player.
It's the age of sail; we're using magnets for direction and best-guesses for bearings.
It wouldn't be right to use the exact bearing (or range).

[Polar coordinate system](https://en.wikipedia.org/wiki/Polar_coordinate_system)


## Compass Rose
I have a concept of North, South, East, and West.
For small scales, we can ignore the curvature of the Earth and assume:
1. 0,1 is North of 0,0
1. 1,0 is East of 0,0
1. 0,-1 is South of 0,0
1. -1, 0 is West of 0,0

That brings up a challenge reporting things to the player.

The program can say that `HMS Speedy` is at (0,0),
with a heading of NNE,
the current is pushing SSE at 2 knots,
the wind is from the SE at 7 knots,
and the `El Gamo` is to the NE at (55,55),
heading W at 8 knots.


# Errata

I am using the 1983 edition of the rulebook.
I have not seen an official set of errata for the game.

This section will note the few typos I've found, plus any odd interpretations of the rules I've made.

# Glossary

See [Glossary of Nautical Terms](https://en.wikipedia.org/wiki/Glossary_of_nautical_terms) on the Wikipedia.

# Notes
1. [Point of Sail](https://en.wikipedia.org/wiki/Point_of_sail)
1. [Effect of point of sail on forces](https://en.wikipedia.org/wiki/Forces_on_sails#Effect_of_points_of_sail_on_forces)

# Links

1. The [Wikipedia page](https://en.wikipedia.org/wiki/Heart_of_Oak:_Naval_Miniatures_for_the_Age_of_Fighting_Sail)
1. Another [Wikipedia page]()
1. [Fantasy Games Unlimited](https://en.wikipedia.org/wiki/Fantasy_Games_Unlimited)
1. [Board Game Geek](https://boardgamegeek.com/boardgame/7779/heart-oak-naval-miniatures-rules-age-fighting-sail) review
1. Math Is Fun article on converting [Polar and Cartesian Coordinates](https://www.mathsisfun.com/polar-cartesian-coordinates.html)
1. [Why do we use radians for polar coordinates rather than degrees?](https://math.stackexchange.com/questions/2362943/why-do-we-use-radians-for-polar-coordinates-rather-than-degrees)
1. [Why Radian Measure Makes Life Easier In Mathematics And Physics](https://qedinsight.wordpress.com/2011/03/14/why-radian-measure-makes-life-easier-in-mathematics-and-physics/)
1. [Compass Rose](https://en.wikipedia.org/wiki/Compass_rose)
1. [Spherical Geometry](https://github.com/golang/geo)
