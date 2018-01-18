package gogo

import (
	"reflect"
	"testing"
)

func Test_17_letterCombinations(t *testing.T) {
	digits := "23"
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	actual := letterCombinations(digits)
	if reflect.DeepEqual(expected, actual) != true {
		t.Fatalf("input: %v, expected: %v, actual: %v", digits, expected, actual)
	}
}
