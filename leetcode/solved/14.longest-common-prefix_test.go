package gogo

import "testing"

func Test_14_longestCommonPrefix(t *testing.T) {
    test1 := []string{}
    expected1 := ""
    actual1 := longestCommonPrefix(test1)
    if expected1 != actual1 {
        t.Fatalf("Empty array tests failed")
    }

    test2 := []string{"hello", "help", "hell",}
    expected2 := "hel"
    actual2 := longestCommonPrefix(test2)
    if expected2 != actual2 {
        t.Fatalf("test: %v, expected: %v, actual: %v", test2, expected2, actual2)
    }

    test3 := []string{"a", "a", "a",}
    expected3 := "a"
    actual3 := longestCommonPrefix(test3)
    if expected3 != actual3 {
        t.Fatalf("test: %v, expected: %v, actual: %v", test3, expected3, actual3)
    }

    test4 := []string{"bbbbbbbbbb","bbbbbbbbbb","bbbbbbbbbb",}
    expected4 := "bbbbbbbbbb"
    actual4 := longestCommonPrefix(test4)
    if expected4 != actual4 {
        t.Fatalf("test: %v, expected: %v, actual: %v", test4, expected4, actual4)
    }
}
