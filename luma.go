package imagehash

import (
	"image"

	"github.com/vitali-fedulov/images4"
)

// lumaVector returns luma values at sample pixels of the icon.
func lumaVector(icon images4.IconT, sample []image.Point) (v []float64) {
	for i := range sample {
		c1, _, _ := images4.Get(icon, images4.IconSize, sample[i])
		v = append(v, float64(c1))
	}
	return v
}
