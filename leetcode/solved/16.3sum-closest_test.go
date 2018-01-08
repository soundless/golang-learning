package gogo

import "testing"

func Test_16_3sum_closest(t *testing.T) {
    S := []int{-1, 2, 1, -4}
    target := 1
    expected := 2
    actual := threeSumClosest(S, target)
    if expected != actual {
        t.Fatalf("tests: %v, target: %v, expected: %v, actual: %v", S, target, expected, actual)
    }
}
