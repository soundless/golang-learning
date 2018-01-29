package gogo

import (
	"testing"
	"reflect"
)

func Test_22_generateParentheses(t *testing.T) {
	input := 3
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	actual := generateParentheses(input)
	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("input: %v, expected: %v, actual: %v", input, expected, actual)
	}
}
