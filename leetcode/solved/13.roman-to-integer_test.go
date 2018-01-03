package gogo

import "testing"

func Test_13_RomanToInt(t *testing.T) {
    tests := map[string]int{
        "I": 1,
        "VIII": 8,
        "X": 10,
        "XXIII": 23,
        "XCIX": 99,
        "CXXIII": 123,
        "DXLV": 545,
        "M": 1000,
        "MXXIV": 1024,
        "MCMXCIX": 1999,
        "MM": 2000,
        "MMMCMXCIX": 3999,
    }

    for key, value := range tests {
        actual := romanToInt(key)
        expected := value
        if expected != actual {
            t.Fatalf("test: %v, expected: %v, actual: %v", key, expected, actual)
        }
    }

    if 0 != romanToInt("") {
        t.Fatalf("Failed to test empty string")
    }
}
