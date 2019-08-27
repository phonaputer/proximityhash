//TODO add full test coverage
package proximityhash

import (
	"github.com/mmcloughlin/geohash"
	"math"
	"testing"
)

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
