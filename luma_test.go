package imagehash

import (
	"image"
	"math"
	"testing"

	"github.com/vitali-fedulov/images4"
)

func TestLumaVector(t *testing.T) {
	icon := images4.EmptyIcon()
	icon.Pixels = make([]uint16, images4.IconSize*images4.IconSize*3)
	expectedColor1 := float64(13.1)
	expectedColor2 := float64(9.1)
	images4.Set(icon, images4.IconSize,
		image.Point{1, 1}, expectedColor1, 29.9, 95.9)
	images4.Set(icon, images4.IconSize,
		image.Point{9, 5}, expectedColor2, 11.0, 12.9)
	got := lumaVector(icon, []image.Point{{1, 1}, {9, 5}})
	if math.Abs(float64(got[0])-expectedColor1) > 0.1 ||
		math.Abs(float64(got[1])-expectedColor2) > 0.1 {
		t.Errorf(
			`Expected 2 color values %v and %v.
			 Got: %v and %v.`, expectedColor1, expectedColor2,
			float64(got[0]), float64(got[1]))
	}
}
