package gogo

import "testing"

func Test_1_2Sum(t *testing.T) {
	nums := []int{7, 11, 2, 15}
	var indexes []int
	indexes = twoSum(nums, 9)
	if indexes[0] != 1 {
		if indexes[1] != 2 {
			t.Fatalf("Failed - index 1: %v, index 2: %v", indexes[0], indexes[1])
		}
	}
}
