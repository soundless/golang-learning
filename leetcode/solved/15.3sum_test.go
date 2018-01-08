package gogo

import (
    "testing"
    "reflect"
)

func Test_15_3_sum(t *testing.T) {
    input := []int{-1, 0, 1, 2, -1, -4}
    expected := [][]int{
        {-1, 0, 1},
        {-1, -1, 2},
    }
    actual := threeSum(input)
    if reflect.DeepEqual(expected, actual) != true {
        t.Fatalf("input: %v, expected: %v, actual: %v", input, expected, actual)
    }
}
