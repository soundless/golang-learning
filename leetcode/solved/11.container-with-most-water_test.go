package gogo

import "testing"

func Test_11_maxArea(t *testing.T) {
	height := []int{1, 2, 3, 4, 3, 2, 1, 5}
	expected := 16
	actual := maxArea(height)
	if expected != actual {
		t.Fatalf("test %v, expected: %v, actual: %v", height, expected, actual)
	}
}
