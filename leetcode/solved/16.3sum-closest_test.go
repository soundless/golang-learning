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

    S1 := []int{-1, 0, 1, 2, -1, -4}
    target1 := 4
    expected1 := 3
    actual1 := threeSumClosest(S1, target1)
    if expected1 != actual1 {
        t.Fatalf("tests: %v, target: %v, expected: %v, actual: %v", S1, target1, expected1, actual1)
    }
}
