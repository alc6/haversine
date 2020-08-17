package haversine

import (
	"math"
)

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
	Lat float64
	Lon float64
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

// radiansToDegrees converts from radians to degrees.
func radiansToDegrees(d float64) float64 {
	return d *  180 / math.Pi
}

// Distance calculates the shortest path between two coordinates on the surface
// of the Earth. This function returns two units of measure, the first is the
// distance in miles, the second is the distance in kilometers.
func Distance(p, q Coord) (mi, km float64) {
	lat1 := degreesToRadians(p.Lat)
	lon1 := degreesToRadians(p.Lon)
	lat2 := degreesToRadians(q.Lat)
	lon2 := degreesToRadians(q.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi = c * earthRadiusMi
	km = c * earthRaidusKm

	return mi, km
}

// Returns the point at given fraction between ‘this’ point and given point.
func (inpoint Coord) IntermediatePointTo(point Coord, fraction float64) (newPt Coord) {
	// https://www.movable-type.co.uk/scripts/latlong.html
	lat1 := degreesToRadians(inpoint.Lat)
	lon1 := degreesToRadians(inpoint.Lon)
	lat2 := degreesToRadians(point.Lat)
	lon2 := degreesToRadians(point.Lon)

	// distance between points
	deltaLat := lat2 - lat1
	deltaLon := lon2 - lon1
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1)*math.Cos(lat2)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	dBetween := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	A := math.Sin((1-fraction)*dBetween) / math.Sin(dBetween)
	B := math.Sin(fraction*dBetween) / math.Sin(dBetween)

	x := A*math.Cos(lat1)*math.Cos(lon1) + B*math.Cos(lat2)*math.Cos(lon2)
	y := A*math.Cos(lat1)*math.Sin(lon1) + B*math.Cos(lat2)*math.Sin(lon2)
	z := A*math.Sin(lat1) + B*math.Sin(lat2)

	lat3Rad := math.Atan2(z, math.Sqrt(x*x+y*y))
	lon3Rad := math.Atan2(y, x)

	newPt.Lat = radiansToDegrees(lat3Rad)
	newPt.Lon = radiansToDegrees(lon3Rad)
	return newPt
}
