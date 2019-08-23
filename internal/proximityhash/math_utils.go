package proximityhash

import (
	"math"
)

const (
	meanEarthRadKm  = 6371.0
	meanEarthDiamKm = meanEarthRadKm * 2
)

// The point struct represents a geographic location specified by a latitude, longitude coordinate pair.
type point struct {
	lat float64
	lng float64
}

func (p *point) equals(p2 point) bool {
	return p.lng == p2.lng && p.lat == p2.lat
}

// toRad converts an angle from degrees to radians.
func toRad(deg float64) float64 {
	return deg * math.Pi / 180
}

func haversinFunc(a float64) float64 {
	return 0.5 * (1 - math.Cos(a))
}

// haversinDist calculates the haversine distance between two lat, lng points.
func haversinDist(x, y point) float64 {
	sqrt := math.Sqrt(haversinFunc(y.lat-x.lat) + math.Cos(x.lat)*math.Cos(y.lat)*haversinFunc(y.lng-x.lng))

	return meanEarthDiamKm * math.Asin(sqrt)
}

// distToLine calculates the shortest distance from a point to a line. The parameters start and end are the start and
// end points of the line.
func distToLine(pt, start, end point) float64 {
	if start.equals(end) {
		return haversinDist(pt, end)
	}

	pLat := toRad(pt.lat)
	pLng := toRad(pt.lng)
	sLat := toRad(start.lat)
	sLng := toRad(start.lng)
	eLat := toRad(end.lat)
	eLng := toRad(end.lng)

	esLat := eLat - sLat
	esLng := eLng - sLng

	u := ((pLat-sLat)*esLat + (pLng-sLng)*esLng) / (esLat*esLat + esLng*esLng)

	if u <= 0 {
		return haversinDist(pt, start)
	}
	if u >= 1 {
		return haversinDist(pt, end)
	}

	sa := point{lat: (pt.lat - start.lat), lng: (pt.lng - start.lng)}
	sb := point{lat: u * (end.lat - start.lat), lng: u * (end.lng - start.lng)}

	return haversinDist(sa, sb)
}

func handleCrossPrimeMerid(lng float64) float64 {
	if lng > 180.0 || lng < -180.0 {
		sign := 1
		if math.Signbit(lng) {
			sign = -1
		}

		return ((lng * sign) - 360.0) * sign
	}

	return lng
}

// addToPoint adds in meters to a given point (the point being specified as a lat, lng pair of angles).
func addToPoint(pt point, deltaLng, deltaLat float64) point {
	latDiff := (deltaLat / meanEarthRadKm) * (180 / math.Pi)
	lngDiff := ((deltaLng / meanEarthRadKm) * (180 / math.Pi)) / math.Cos(pt.lat*(math.Pi/180))

	return point{
		lat: pt.lat + latDiff,
		lng: handleCrossPrimeMerid(pt.lng + lngDiff),
	}
}
