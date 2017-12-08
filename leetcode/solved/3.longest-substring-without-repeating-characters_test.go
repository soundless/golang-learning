package gogo

import "testing"

func Test_longestSubstringWithoutRepeatingCharacters(t *testing.T) {
	strs := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}
	for key, value := range strs {
		if lengthOfLongestSubstring(key) != value {
			t.Fatalf("Failed")
		}
	}
}
