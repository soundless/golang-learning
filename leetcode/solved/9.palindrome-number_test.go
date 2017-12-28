package gogo

import "testing"

func Test_9_isPalindrome(t *testing.T) {
    tests := []int{12321, 1221, -1, -12321, 0}
    results := []bool{true, true, false, false, true}

    for i := 0; i < len(tests); i++ {
        result := isPalindrome(tests[i])
        if result != results[i] {
            t.Fatalf("test: %v, expected: %v, actual: %v", tests[i], results[i], result)
        }
    }
}
