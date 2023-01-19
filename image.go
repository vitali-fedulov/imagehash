package imagehash

import (
	"fmt"
	"image"

	"github.com/vitali-fedulov/images4"
)

// Default hyper space parameters.
const (
	epsPct     = 0.25
	numBuckets = 4
)

// HyperParams holds hyper space parameters for hash computations.
type HyperParams struct {
	epsPct     float64
	numBuckets int
}

// Open opens and decodes an image file for a given path. It is a convenience wrapper around
// imagehash and images4 that returns the icon and hashes.
//
// Takes as input a single optional params that holds epsPct and numBuckets hyper space parameters.
func Open(path string, params ...HyperParams) (Image, error) {
	if len(params) > 1 {
		return Image{}, fmt.Errorf("Open() only accepts a single HyperParams as input, got %d instead", len(params))
	}
	eps := epsPct
	buckets := numBuckets
	if len(params) == 1 {
		eps = params[0].epsPct
		buckets = params[0].numBuckets
	}

	img, err := images4.Open(path)
	if err != nil {
		return Image{}, err
	}
	icon := images4.Icon(img)
	centralHash := CentralHash(icon, HyperPoints10, eps, buckets)
	hashSet := HashSet(icon, HyperPoints10, eps, buckets)
	return Image{Image: img, Icon: icon, CentralHash: centralHash, HashSet: hashSet}, nil
}

// Image is a convenience wrapper for holding everything imagehash and images4 needs.
//
// Call Similar() on an Image object for quick similarity computation instead of images4.Similar().
type Image struct {
	Image       image.Image
	Icon        images4.IconT
	CentralHash uint64
	HashSet     []uint64
}

// Similar returns true if img2 is similar to img. It is a convenience wrapper around imagehash and
// images4 that first compares the images using CentralHash & Hashset, and then using the
// images4.Similar() function iff there is a match between the hashes.
//
// It is faster than calling images4.Similar() directly for dissimilar images and
// only slightly slower for similar images.
func (img Image) Similar(img2 Image) bool {
	foundSimilarImage := false
	for _, hash := range img2.HashSet {
		if img.CentralHash == hash {
			foundSimilarImage = true
			break
		}
	}
	return foundSimilarImage && images4.Similar(img.Icon, img2.Icon)
}
