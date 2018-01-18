package gogo

import "testing"

func Test_20_isValid(t *testing.T) {
	valid := []string{"()", "()[]{}", "[()]{}"}
	invalid := []string{"(]", "([)]", ")"}

	for _, s := range valid {
		expected := true
		actual := isValid(s)
		if expected != actual {
			t.Fatalf("input: %v, expected: %v, actual: %v", s, expected, actual)
		}
	}

	for _, s := range invalid {
		expected := false
		actual := isValid(s)
		if expected != actual {
			t.Fatalf("input: %v, expected: %v, actual: %v", s, expected, actual)
		}
	}
}
