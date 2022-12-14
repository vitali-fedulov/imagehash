package imagehash

import (
	"image"
	"reflect"
	"testing"

	"github.com/vitali-fedulov/images4"
)

var p1 = []uint16{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255}

var p2 = []uint16{
	231, 183, 148, 21, 47, 16, 69, 45, 151, 64, 181,
	49, 243, 60, 147, 177, 19, 85, 205, 146, 236,
	67, 28, 163, 143, 94, 138, 214, 149, 103, 57,
	145, 74, 249, 235, 212, 30, 202, 88, 244, 196,
	118, 95, 59, 148, 49, 165, 92, 10, 39, 253, 16,
	2, 166, 126, 196, 22, 53, 84, 156, 117, 178, 8,
	171, 134, 189, 7, 226, 156, 51, 211, 230, 14,
	242, 165, 4, 246, 208, 35, 176, 248, 91, 81,
	174, 68, 166, 15, 199, 78, 233, 75, 99, 117,
	55, 140, 194, 127, 133, 198, 173, 2, 126, 84,
	57, 88, 150, 72, 168, 41, 44, 43, 169, 143, 96,
	75, 33, 204, 161, 153, 120, 178, 99, 181, 49,
	132, 91, 38, 164, 45, 87, 92, 80, 85, 87, 195,
	18, 220, 155, 66, 145, 239, 61, 222, 206, 235,
	141, 75, 247, 223, 102, 232, 163, 17, 213, 221,
	161, 175, 27, 191, 178, 243, 220, 86, 149, 34,
	85, 185, 140, 30, 171, 213, 244, 192, 79, 254,
	211, 9, 167, 176, 246, 62, 245, 42, 29, 192,
	12, 10, 250, 10, 89, 182, 219, 49, 51, 38, 38,
	25, 21, 76, 182, 163, 22, 200, 183, 242, 157,
	71, 57, 99, 102, 84, 183, 143, 239, 220, 237,
	221, 161, 167, 92, 245, 42, 223, 104, 50, 0,
	106, 131, 53, 56, 177, 208, 49, 125, 212, 255,
	169, 126, 154, 240, 161, 61, 23, 163, 60, 176,
	138, 250, 211, 205, 134, 198, 36, 212, 206,
	29, 233, 107, 109, 104, 36, 81, 53, 167, 204,
	25, 60, 143, 234, 246, 63, 181, 109, 100, 149,
	28, 56, 195, 155, 114, 141, 70, 139, 222, 207,
	30, 213, 154, 238, 98, 117, 96, 182, 85, 17,
	241, 21, 250, 4, 241, 173, 217, 53, 169, 158,
	252, 221, 237, 252, 85, 59, 25, 209, 191, 120,
	191, 122, 228, 240, 254, 133, 74, 190, 59, 180,
	220, 27, 90, 197, 183, 179, 164, 188, 52, 105,
	8, 149, 248, 11, 132, 248, 126, 99, 118, 227,
	163, 82, 236, 98, 206, 218, 154, 231, 5, 201,
	125, 47, 238, 152, 154, 12, 115, 225}

var p3 = []uint16{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0}

func TestHashSet(t *testing.T) {
	want := []uint64{
		13332021, 1013332021, 13332121,
		1013332121, 13332022, 1013332022,
		13332122, 1013332122}
	icon2 := images4.EmptyIcon()
	for i := range p2 {
		p2[i] = 255 * p2[i] // Premultiplication.
	}
	icon2.Pixels = p2
	got := HashSet(icon2, HyperPoints10, 0.25, 4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v.", want, got)
	}

	want = []uint64{0}
	icon3 := images4.EmptyIcon()
	for i := range p3 {
		p3[i] = 255 * p3[i] // Premultiplication.
	}
	icon3.Pixels = p3
	got = HashSet(icon3, HyperPoints10, 0.25, 4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v.", want, got)
	}
}

func TestCentralHash(t *testing.T) {
	want := uint64(3333333333)
	icon1 := images4.EmptyIcon()
	for i := range p1 {
		p1[i] = 255 * p1[i] // Premultiplication.
	}
	icon1.Pixels = p1
	got := CentralHash(icon1, HyperPoints10, 0.25, 4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v.", want, got)
	}

	want = uint64(13332021)
	icon2 := images4.EmptyIcon()
	icon2.Pixels = p2
	got = CentralHash(icon2, HyperPoints10, 0.25, 4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v, got %v.", want, got)
	}

}

func TestHyperPoints10(t *testing.T) {
	s := HyperPoints10
	if len(s) != 10 {
		t.Errorf(
			"Mismatching set length. Expected 10, got %v",
			len(s))
	}
	for _, p1 := range s {
		for _, p2 := range s {
			if p1 == p2 {
				continue
			}
			d := distance(p1, p2)
			if d < 2 {
				t.Errorf(
					`Too small distance between points, got
					distance %v, between points %v and %v`,
					d, p1, p2)
			}
		}
	}
}

// The test can be brittle, but should be fine, if you rerun it a few times,
// because point selection is not strict in terms of algorithm performance.
func TestSelectPoints(t *testing.T) {
	got := CustomPoints(5)
	want := map[image.Point]bool{
		{2, 2}: true, {2, 8}: true, {5, 5}: true,
		{8, 2}: true, {8, 8}: true}
	for k := range got {
		if _, ok := want[k]; !ok {
			t.Errorf(
				"Missing point %v in expected %v", k, want)
		}
	}
	got = CustomPoints(12)
	if len(got) != 12 {
		t.Errorf(
			"Mismatching set length. Expected 12, got %v",
			len(got))
	}
	if _, ok := got[image.Point{1, 1}]; !ok {
		t.Errorf(
			"Missing point{1, 1} in the result %v", got)
	}
	for p1 := range got {
		for p2 := range got {
			if p1 == p2 {
				continue
			}
			d := distance(p1, p2)
			if d < 2 {
				t.Errorf(
					`Too small distance between points, got
					distance %v, between points %v and %v`,
					d, p1, p2)
			}
		}
	}
}

func TestDistance(t *testing.T) {
	got := distance(
		image.Point{5, 7}, image.Point{2, 8})
	want := 3.1622776601683795
	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestMinKey(t *testing.T) {
	got := minKey(
		map[image.Point]float64{
			{1, 1}: 1.9, {2, 2}: 0.3, {3, 3}: 0.01,
			{7, 7}: 12.0, {9, 9}: 3.0})
	want := image.Point{3, 3}
	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestMaxKey(t *testing.T) {
	got := maxKey(
		map[image.Point]float64{
			{1, 1}: 1.9, {2, 2}: 0.3, {3, 3}: 0.01,
			{7, 7}: 12.0, {9, 9}: 3.0})
	want := image.Point{7, 7}
	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestExlude(t *testing.T) {
	got := exclude(
		image.Point{2, 2}, map[image.Point]bool{
			{1, 1}: true, {2, 2}: true, {3, 3}: true,
			{7, 7}: true, {9, 9}: true})
	want := map[image.Point]bool{
		{1, 1}: true, {3, 3}: true,
		{7, 7}: true, {9, 9}: true}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestNearest(t *testing.T) {
	got := nearest(
		map[image.Point]bool{
			{1, 1}: true, {2, 2}: true, {3, 3}: true,
			{7, 7}: true, {9, 9}: true},
		image.Point{6, 6})
	want := image.Point{7, 7}
	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
	got = nearest(
		map[image.Point]bool{
			{1, 1}: true, {2, 2}: true, {3, 3}: true,
			{7, 7}: true, {9, 9}: true},
		image.Point{3, 3})
	want = image.Point{2, 2}
	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}
