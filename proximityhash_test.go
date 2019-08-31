package proximityhash

import (
	"github.com/mmcloughlin/geohash"
	"math"
	"testing"
)

func Test_FindGeohashesWithinRadius_radiusDoesNotIntersectSidesOfAnyGeohash_shouldReturnOnlyHashOfInputPoint(t *testing.T) {
	precision := uint(12)
	buckinghamFountainGeohash := "dp3wnpmb4ekt";
	lat, lng := geohash.DecodeCenter(buckinghamFountainGeohash)
	radius := 0.00000001

	resFll, resPrt := FindGeohashesWithinRadius(lat,lng, radius, precision)

	assertSize(t, resFll, 0)
	assertSize(t, resPrt, 1)
	assertSliceContainsOtherSlice(t, []string{buckinghamFountainGeohash}, resPrt)
}

func Test_FindGeoHashesWithinRadius_radiusIsSmallAndCenterPointLiesOnBoundaryOfTwoGeohashes_shouldReturnTwoBoundaryHashes(t *testing.T) {
	precision := uint(9)
	bascomHillGeohash := "dp8mj9e71"
	rect := geohash.BoundingBox(bascomHillGeohash)
	lat := rect.MaxLat - 0.00003
	lng := rect.MinLng
	radius := 0.1

	resFll, resPrt := FindGeohashesWithinRadius(lat,lng, radius, precision)

	assertSize(t, resFll, 0)
	assertSize(t, resPrt, 2)
	assertSliceContainsOtherSlice(t, []string{bascomHillGeohash, "dp8mj9e70"}, resPrt)
}

func Test_FindGeohashesWithinRadius_radiusDoesntFullyIncludeParentHashButTouchesAllNeighbors_centerAndNSEWneighborsArePartial(t *testing.T) {
	precision := uint(5)
	sapporoHash := "xpssb"
	lat, lng := geohash.DecodeCenter(sapporoHash)
	halfWidth := 4890.0 / 2
	radius := halfWidth + 1

	resFll, resPrt := FindGeohashesWithinRadius(lat, lng, radius, precision)

	assertSize(t, resFll, 0)
	assertSize(t, resPrt, 5)
	assertSliceContainsOtherSlice(t, []string{"xpst0", "xpssc", "xpss8", "xpskz", sapporoHash}, resPrt)
}

func Test_FindGeohashesWithinRadius_radiusBarelyLargeEnoughToFullMatchCenterHash_centerFull8neighborsPartial(
	t *testing.T) {
	precision := uint(7)
	shibuyaHash := "xn76fgr"
	lat, lng := geohash.DecodeCenter(shibuyaHash)
	halfWidth := 153.0 / 2
	radius := math.Sqrt(2 * math.Pow(halfWidth, 2))

	resFll, resPrt := FindGeohashesWithinRadius(lat, lng, radius, precision)

	assertSize(t, resFll, 1)
	assertSliceContainsOtherSlice(t, []string{"xn76fgr"}, resFll)
	assertSize(t, resPrt, 8)
	assertSliceContainsOtherSlice(t, []string{"xn76g50", "xn76fgp", "xn76fgn", "xn76fgq", "xn76fgw", "xn76fgx",
		"xn76g58", "xn76g52"}, resPrt)
}

func Test_FindGeohashesWithinRadius_radiusLargeEnoughToFullMatchFourHashes(t *testing.T) {
	precision := uint(3)
	uluruHash := "qgm"
	rect := geohash.BoundingBox(uluruHash)
	radius := 1.5 * 156000

	resFll, resPrt := FindGeohashesWithinRadius(rect.MaxLat, rect.MinLng, radius, precision)

	assertSize(t, resFll, 4)
	assertSliceContainsOtherSlice(t, []string{uluruHash, "qgt", "qgs", "qgk"}, resFll)
	assertSize(t, resPrt, 12)
	assertSliceContainsOtherSlice(t, []string{"qgw", "qg7", "qgg", "qgv", "qgy", "qgh", "qgj", "qgn", "qgq", "qge",
		"qgu", "qg5"}, resPrt)
}

func Test_FindGeohashesWithinRadius_radiusLargeEnoughToFullMatchManyHashes(t *testing.T) {
	precision := uint(2)
	balkansItalyHash := "sr"
	rect := geohash.BoundingBox(balkansItalyHash)
	radius := 3.0 * 1250000.0

	resFll, resPrt := FindGeohashesWithinRadius(rect.MaxLat, rect.MinLng, radius, precision)

	assertSize(t, resFll, 64)
	assertSliceContainsOtherSlice(t, []string{balkansItalyHash, "ub", "uc", "ud", "ue", "uf", "ug", "uh", "uk",
		"us", "uu", "s5", "s7", "g2", "g3", "er", "et", "g6", "eu", "g7", "g8", "ev", "g9", "ew", "ex", "ey",
		"ez", "sh", "sj", "sk", "sm", "u0", "gb", "sn", "u1", "gc", "u2", "sp", "gd", "u3", "sq", "ge", "u4",
		"gf", "u5", "gg", "ss", "u6", "st", "u7", "u8", "sv", "u9", "sw", "sx", "sy", "sz", "gs", "gu", "v0",
		"v1", "tp", "v4", "v5"}, resFll)
	assertSize(t, resPrt, 40)
	assertSliceContainsOtherSlice(t, []string{"tn", "tj", "vh", "vk", "v7", "v6", "v3", "v2", "tr", "tq", "th",
		"su", "es", "ek", "em", "eq", "gk", "um", "ut", "uj", "uv", "gv", "vj", "gt", "gm", "sg", "se", "eg",
		"ee", "en", "ep", "g0", "g1", "g4", "g5", "gh", "sd", "s6", "s4", "ef"}, resPrt)
}


/*
I am leaving in this commented, failing test to illustrate that this library is not guaranteed to work correctly near
the poles. And, in fact, it does not work correctly near the poles.
 */

// FAILING
//func Test_FindGeohashesWithinRadius_radiusBarelyLargeEnoughToFullMatchCenterHash_antarctica(t *testing.T) {
//	precision := uint(7)
//	antarcticaHash := "h915"
//	lat, lng := geohash.DecodeCenter(antarcticaHash)
//	halfWidth := 153.0 / 2
//	radius := math.Sqrt(2 * math.Pow(halfWidth, 2))
//
//	resFll, resPrt := FindGeohashesWithinRadius(lat, lng, radius, precision)
//
//	assertSize(t, resFll, 1)
//	assertSliceContainsOtherSlice(t, []string{antarcticaHash}, resFll)
//	assertSize(t, resPrt, 8)
//	assertSliceContainsOtherSlice(t, []string{"h90u", "h91h", "h91k", "h917", "h916", "h914", "h90f", "h90g"},
//		resPrt)
//}

func assertSize(t *testing.T, slice []string, expectedSize int) {
	realSize := len(slice)
	if realSize != expectedSize {
		t.Fatalf("incorrect size! expected '%v' but got '%v'", expectedSize, realSize)
	}
}

func assertSliceContainsOtherSlice(t *testing.T, expected, real []string) {
	realHasItem := make(map[string]bool)
	for _, item := range real {
		realHasItem[item] = true
	}

	for _, item := range expected {
		if !realHasItem[item] {
			t.Fatalf("slice didn't contain expected element! expected to find: %v", item)
		}
	}
}
