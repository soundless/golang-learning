package gogo

import "testing"

func Test_12_intToRoman(t *testing.T) {
	tests := map[int]string{
		1:    "I",
		8:    "VIII",
		10:   "X",
		23:   "XXIII",
		99:   "XCIX",
		123:  "CXXIII",
		545:  "DXLV",
		1000: "M",
		1024: "MXXIV",
		1999: "MCMXCIX",
		2000: "MM",
		3999: "MMMCMXCIX",
	}

	for key, value := range tests {
		actual := intToRoman(key)
		if actual != value {
			t.Fatalf("test: %v, expected: %v, actual: %v", key, value, actual)
		}
	}
}
