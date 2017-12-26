package gogo

import "testing"

func Test_reverse(t *testing.T) {
    tests := []int{321, -123, 1234567899, -1234567899, 0, 100, 10000}
    results := []int{123, -321, 0, 0, 0, 1, 1}

    for i := 0; i < len(tests); i++ {
        if result := reverse(tests[i]); result != results[i] {
            t.Fatalf("expected: %v, actual: %v", results[i], result)
        }
    }
}
