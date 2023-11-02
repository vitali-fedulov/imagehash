# Large scale image search with Go

This hash-table-based package provides fast and RAM-friendly **rough approximation** of image similarity for VERY LARGE image collections (millions and more). Another use-case is for computers with little memory or computing power.

The package is a pre-filtering **first step**. The **second step** requires a more precise and slower package [images4](https://github.com/vitali-fedulov/images4) on the image set produced in the first step. This 2 step sequence (imagehash > images4) is necessary, because direct one-to-all comparison with images4 might be slow for very large image collections.

**Recommendation:** Try to use [images4](https://github.com/vitali-fedulov/images4) directly, and ONLY IF it searches slowly in your image set, THEN you need this package. *(Just do not forget to save your icons in images4, so you do not have to perform the very slow image file decoding again - this is the purpose of the icons).*

---

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/imagehash) for code reference

[Algorithm](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html)

Note for curious users: *An alternative to using images4 package is generating multiple hash sets on different pixel sub-sets of the icon with package imagehash, so that search results of one hash set can be joined with another, or several hash sets. Each join operation will improve the result. Look at var `HyperPoints10` description to understand how to create such different pixel sub-sets.*

---

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
