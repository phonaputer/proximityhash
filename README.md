## ProximityHash Go: Geohashes in Proximity to a Point
* *Ported from the [ProximityHash project](https://github.com/ashwin711/proximityhash) created by ashwin711*

[Geohash](https://en.wikipedia.org/wiki/Geohash) is a geocoding system invented by Gustavo Niemeyer and placed into the
public domain. It is a hierarchical spatial data structure which subdivides space
into buckets of grid shape, which is one of the many applications of what is known
as a Z-order curve, and generally space-filling curves.

### ProximityHash Go
ProximityHash generates a set of geohashes that cover a circular area, given the
center coordinates and the radius. It returns two sets. One contains all geohashes
falling entirely within the radius. The other contains all geohashes which fall
partially within the radius. On a map, the results would look like a (rough) circle of
geohashes which fully match (with the origin point lying at the circle's center)
with a ring of partially matching geohashes around the perimeter of the circle.

### Usage and Documentation
You can download this package using go get: 

```
go get github.com/phonaputer/proximityhash-go
```

and then import it into your Go project like so:

```
import "github.com/phonaputer/proximityhash-go"
```

### Contributors
phonaputer [https://github.com/phonaputer]

### Idea Taken From
ProximityHash by Ashwin Nair [https://github.com/ashwin711/proximityhash]
