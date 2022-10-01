# Fast approximate image search with Go

This hash-table-based package provides fast and RAM-friendly **rough approximation** of image similarity. It is meant for datasets with billions images or more (not small image sets) as a pre-filtering step. For small image sets use package [images4](https://github.com/vitali-fedulov/images4) instead.

**Important next step** is using a slower one-to-all image comparison package [images4](https://github.com/vitali-fedulov/images4) on the resulting image set produced with package imagehash. This 2 step sequence (imagehash > images4) is necessary, because direct one-to-all comparison with images4 is not scalable for very large image sets.

An alternative to using images4 package are generating multiple hash sets on different pixel sub-sets of the icon, so that search results of one hash set can be joined with another, or several hash sets. Each join operation will improve the result. Look at var `HyperPoints10` description to understand how to create such different pixel sub-sets.

[Algorithm](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html)

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/imagehash) for code reference.

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
	epsPct = 0.25
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
