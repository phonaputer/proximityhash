package proximityhash

import (
	"testing"
)

var Test_PointEquals_pointsAreEqual_returnsTrue_cases = [] struct {
	name string
	a point
	b point
} {
	{ "simple equality", point{1.1, 2.2}, point{1.1, 2.2} },
	{ "lng not set", point{lat: 1.1}, point{lat: 1.1} },
	{ "lat not set", point{lng: 1.1}, point{lng: 1.1} },
	{ "both not set", point{}, point{} },
}
func Test_PointEquals_pointsAreEqual_returnsTrue(t *testing.T) {
	for _, cse := range Test_PointEquals_pointsAreEqual_returnsTrue_cases {
		res := cse.a.equals(cse.b)

		if !res {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}

var Test_PointEquals_pointsAreNotEqual_returnsFalse_cases = [] struct {
	name string
	a point
	b point
} {
	{ "lat not equal", point{1.1, 2.2}, point{1.0, 2.2} },
	{ "lng not equal", point{1.1, 2.2}, point{1.1, 2.1} },
	{ "both not equal", point{1.1, 2.2}, point{1.0, 2.1} },
	{ "lng not set", point{1.1, 2.2}, point{lat: 1.1} },
	{ "lat not set", point{1.1, 2.2}, point{lng: 2.2} },
	{ "both not set", point{1.1, 2.2}, point{} },
}
func Test_PointEquals_pointsAreNotEqual_returnsFalse(t *testing.T) {
	for _, cse := range Test_PointEquals_pointsAreNotEqual_returnsFalse_cases {
		res := cse.a.equals(cse.b)

		if res {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}

var Test_toRad_cases = [] struct {
	name string
	deg float64
	expctdRad float64
} {
	{"0 deg", 0, 0},
	{"45 deg", 45, 0.7853981633974483},
	{"90 deg", 90, 1.5707963267948966},
	{"135 deg", 135, 2.356194490192345},
	{"180 deg", 180, 3.141592653589793},
}
func Test_toRad(t *testing.T){
	for _, cse := range Test_toRad_cases {
		res := toRad(cse.deg)

		if cse.expctdRad != res {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}

var Test_haversinDist_cases = [] struct {
	name string
	a point
	b point
	expctdOut float64
} {
	{ "japan short", point{35.544113, 139.516296}, point{35.549344, 139.516184}, 0.581748916353381 },
	{ "NE hemisphere to NW hemisphere", point{35.544113, 139.516296}, point{41.396292, -87.943309},
		10187.672469669049},
	{ "NE hemisphere to SW hemisphere", point{35.544113, 139.516296}, point{-54.943748, -66.816214},
		17065.499246347797},
	{ "south polar region", point{-88.482081, 3.100664}, point{-84.418740, 137.766579}, 748.9113590535179},
}
func Test_haversinDist(t *testing.T) {
	for _, cse := range Test_haversinDist_cases {
		res := haversinDist(cse.a, cse.b)

		if cse.expctdOut != res {
			t.Fatalf("failed case: %v", res)
		}
	}
}

var Test_distToLine_cases = [] struct {
	name string
	pt point
	start point
	end point
	expctd float64
} {
	{"origin to line", point{0.0, 0.0}, point{1.0, 1.0}, point{-1.0, 1.0}, 111.17799068882648},
	{"origin to single point", point{0.0, 0.0}, point{1.0, 1.0}, point{1.0, 1.0}, 157.24938127194397},
	{"polar region medium distance", point{-88.482081, 3.100664}, point{-85.418740, 137.766579},
		point{-85.418740, 140.766579}, 639.4089261622778},
	{"japan short distance", point{35.544113, 139.516296}, point{35.549344, 139.516184},
		point{35.54777, 139.51777}, 0.42794904815548024},
}
func Test_distToLine(t *testing.T) {
	for _, cse := range Test_distToLine_cases {
		res := distToLine(cse.pt, cse.start, cse.end)

		if cse.expctd != res {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}

var Test_handleCrossPrimeMerid_cases = [] struct {
	name string
	lng float64
	expctd float64
} {
	{"doesn't cross", 180, 180},
	{"crosses positive", 180.00001, -179.99999},
	{"crosses negative", -180.00001, 179.99999},
	{"crosses a lot positive", 359.0, -1.0},
	{"crosses a lot negative", -359.0, 1.0},

}
func Test_handleCrossPrimeMerid(t *testing.T) {
	for _, cse := range Test_handleCrossPrimeMerid_cases {
		res := handleCrossPrimeMerid(cse.lng)

		if cse.expctd != res {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}

var Test_addToPoint_cases = [] struct {
	name string
	pt point
	dLat float64
	dLng float64
	expctd point
} {
	{"japan short", point{35.544113, 139.516296}, 100, 100, point{36.44343460591873, 140.62156424524923}},
	{"polar negative", point{-88.482081, 3.100664}, -500, -500, point{-92.97868902959365, -166.64937026798168}},
	{"crosses prime merid", point{-59.482081, 179.999}, 1000, 1000, point{-50.4888649408127, -162.29111861773117}},
	{"crosses prime merid negative", point{-59.482081, -179.999}, -1000, -1000,
		point{-68.47529705918731, 162.29111861773117}},
}
func Test_addToPoint(t *testing.T) {
	for _, cse := range Test_addToPoint_cases {
		res := addToPoint(cse.pt, cse.dLng, cse.dLat)

		if !cse.expctd.equals(res) {
			t.Fatalf("failed case: %v", cse.name)
		}
	}
}