// Proximityhash is a utility which finds all geohashes that lie within a given distance (in meters) of a given
// latitude, longitude coordinate point.
package proximityhash

import (
	"github.com/mmcloughlin/geohash"
)

func checkCircleIntersectsRectangleGeometrically(radius float64, center point, rect geohash.Box) bool {
	nw := point{rect.MaxLat, rect.MinLng}
	ne := point{rect.MaxLat, rect.MaxLng}
	se := point{rect.MinLat, rect.MaxLng}
	sw := point{rect.MinLat, rect.MinLng}

	if radius >= distToLine(center, nw, ne) ||
		radius >= distToLine(center, sw, se) ||
		radius >= distToLine(center, sw, nw) ||
		radius >= distToLine(center, se, ne) {

		return true
	}

	return false
}

func checkInsideRadiusSimple(radius float64, center point, rect geohash.Box) (partial bool, full bool) {
	nwCornerInside := radius >= haversinDist(center, point{rect.MaxLat, rect.MinLng})
	neCornerInside := radius >= haversinDist(center, point{rect.MaxLat, rect.MaxLng})
	seCornerInside := radius >= haversinDist(center, point{rect.MinLat, rect.MaxLng})
	swCornerInside := radius >= haversinDist(center, point{rect.MinLat, rect.MinLng})

	if nwCornerInside && neCornerInside && seCornerInside && swCornerInside {
		return false, true
	}

	if nwCornerInside || neCornerInside || seCornerInside || swCornerInside {
		return true, false
	}

	return false, false
}

func isGeohashInsideRadius(radius float64, center point, hash string) (partial bool, full bool) {
	rect := geohash.BoundingBox(hash)

	prt, fll := checkInsideRadiusSimple(radius, center, rect)

	if prt || fll {
		return prt, fll
	}

	prt = checkCircleIntersectsRectangleGeometrically(radius, center, rect)

	return prt, false
}

func filterAlreadyChecked(toFilter []string, alreadyChecked map[string]bool) []string {
	var res []string

	for _, item := range toFilter {
		if !alreadyChecked[item] {
			res = append(res, item)
			alreadyChecked[item] = true
		}
	}

	return res
}

// The FindGeohashesWithinRadius function finds all geohashes within the given radius of the given (lat, lng)
// coordinate point. The geohashes will have the given precision. Geohashes which are 100% inside the radius will be in
// the fullMatches return value. Geohashes which lie partially but not fully within the radius will be in the
// partialMatches return value.
func FindGeohashesWithinRadius(lat, lng, radius float64, precision uint) (fullMatches, partialMatches []string) {
	alreadyChecked := make(map[string]bool)
	queue := newStringQueue()

	firstHash := geohash.EncodeWithPrecision(lat, lng, precision)
	center := point{lat, lng}

	alreadyChecked[firstHash] = true
	queue.enqueue(firstHash)

	for !queue.isEmpty() {
		curHash, _ := queue.dequeue()
		prt, fll := isGeohashInsideRadius(radius, center, curHash)

		if fll {
			fullMatches = append(fullMatches, curHash)
		} else if prt {
			partialMatches = append(partialMatches, curHash)
		}

		if prt || fll {
			neighbors := geohash.Neighbors(curHash)
			unchecked := filterAlreadyChecked(neighbors, alreadyChecked)

			queue.enqueue(unchecked...)
		}
	}

	if len(fullMatches) < 1 && len(partialMatches) < 1 {
		partialMatches = append(partialMatches, firstHash)
	}

	return fullMatches, partialMatches
}
