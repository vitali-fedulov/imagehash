# Large scale image similarity search (Golang)

This package is a fast and RAM-friendly hash-table-based image comparison tool **for large image collections** (thousands and more). Resized and near-duplicate images can be found with it.

When numBuckets parameter is low (~4), the package is a **rough pre-filtering first step**. Then the **second precise step** is needed with [images4](https://github.com/vitali-fedulov/images4) on the image set produced in the first step. This 2 step sequence (imagehash > images4) is necessary, because direct one-to-all comparison with images4 might be slow for very large image collections. For small image sets it is OK to skip the first step altogether.

When numBuckets is very high (~200), be sure to do a few tests, because, as in the example below, only 10 dimensions (pixels in Y channel) are used from the total 11x11*3=363 pixel values in the icon. This could under-represent some images, because bucket width is very small for high numBuckets.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/imagehash)

[Algorithm](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html)

## Example of comparing 2 photos using imagehash

The demo shows only the hash-based similarity testing (without making actual hash table). But hash table is implied in full implementation.

```go
package main

import (
	"fmt"
	"github.com/vitali-fedulov/imagehash"
	"github.com/vitali-fedulov/images4"
)

const (
	// Recommended hyper-space parameters for initial trials.

	// I usually do not change epsPct parameter.
	// epsPct defines the range of uncertainty at hypercube borders,
	// when a nearest similar point may end up in the nearby hypercube,
	// thus having a different hash. The larger the value, the larger
	// the uncertainty range is. Larger values may produce larger hashSets,
	// which could be compute-expensive. 0.25 corresponds to 25% of bucket
	// width.
	epsPct = 0.25

	// Experiment by increasing numBuckets from 4 to 230 or higher.
	// It will make your searches faster, more precise, but maybe too strict.
	// It corresponds to the level of granularity of hyperspace quantization.
	// The higher the value, the more granular is N-space sub-division.
	// This example uses 10-dimensional vectors, splitting the 10-space into
	// 4^10 = 1048576 hypercubes. 4 splits one pixel brightness values into
	// 4 buckets. For numBuckets = 230, there will be 4×10²³ possible hypercubes.
	numBuckets = 4
)

func main() {

	// Paths to photos.
	path1 := "1.jpg"
	path2 := "2.jpg"

	// Open photos (skipping error handling for clarity).
	img1, _ := images4.Open(path1)
	img2, _ := images4.Open(path2)

	// Icons are compact image representations needed for comparison.
	icon1 := images4.Icon(img1)
	icon2 := images4.Icon(img2)

	// Hash table values.

	// Value to save to the hash table as a key with corresponding
	// image ids. Table structure: map[centralHash][]imageId.
	// imageId is simply an image number in a directory tree.
	centralHash := imagehash.CentralHash(
		icon1, imagehash.HyperPoints10, epsPct, numBuckets)

	// Hash set to be used as a query to the hash table. Each hash from
	// the hashSet has to be checked against the hash table.
	// See more info in the package "hyper" README.
	hashSet := imagehash.HashSet(
		icon2, imagehash.HyperPoints10, epsPct, numBuckets)

	// Checking hash matches. In full implementation this will
	// be done on the hash table map[centralHash][]imageId.
	foundSimilarImage := false
	for _, hash := range hashSet {
		if centralHash == hash {
			foundSimilarImage = true
			break
		}
	}

	// Image comparison result.
	if foundSimilarImage {
		fmt.Println("Images are approximately similar.")
	} else {
		fmt.Println("Images are distinct.")
	}

	// Then use func Similar of package images4 for final
	// confirmation of image similarity. That is:
	// if images4.Similar(icon1, icon2) == true {
	//    fmt.Println("Images are definitely similar")
	// }
}
```

## For advanced users
An alternative to using images4 package is generating multiple hash sets on different pixel sub-sets of the icon with package imagehash, so that search results of one hash set can be joined with another, or several hash sets. Each join operation will improve the result. Look at var `HyperPoints10` description to understand how to create such different pixel sub-sets.
