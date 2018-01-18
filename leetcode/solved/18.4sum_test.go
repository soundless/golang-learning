package gogo

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	expected := [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	actual := fourSum(input, 0)
	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("input: %v, expected: %v, actual: %v", input, expected, actual)
	}
}
