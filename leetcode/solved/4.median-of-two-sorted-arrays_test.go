package gogo

import "testing"

func Test_4_findMedianSortedArray(t *testing.T) {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	r1 := findMedianSortedArrays(a1, a2)
	if r1 != 2.5 {
		t.Fatalf("Result is not %v but %v.", 2.5, r1)
	}

	a3 := []int{1, 3}
	a4 := []int{2}
	r2 := findMedianSortedArrays(a3, a4)
	if r2 != 2.0 {
		t.Fatalf("Result is not %v but %v.", 2.0, r2)
	}
}
